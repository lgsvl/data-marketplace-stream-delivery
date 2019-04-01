package controllers

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/lgsvl/data-marketplace-stream-delivery/resources"
	"github.com/lgsvl/data-marketplace-stream-delivery/utils"
)

type ppDataStreamController struct {
	ctx              context.Context
	logger           *log.Logger
	registry         resources.Registry
	blockchainClient utils.BlockchainClient
	streamHelper     utils.StreamHelper
}

// NewPPDataStreamController creates  new instance of the ppDataStreamController
func NewPPDataStreamController(c context.Context, l *log.Logger, r resources.Registry) resources.DataStreamController {
	return &ppDataStreamController{
		ctx:              c,
		logger:           l,
		registry:         r,
		blockchainClient: utils.NewHyperledgerClient(c, l),
		streamHelper:     utils.NewStreamHelper(c, l)}
}

// NewPPDataStreamControllerWithClients creates  new instance of the ppDataStreamController
func NewPPDataStreamControllerWithClients(c context.Context, l *log.Logger, r resources.Registry, client utils.BlockchainClient, h utils.StreamHelper) resources.DataStreamController {
	return &ppDataStreamController{
		ctx:              c,
		logger:           l,
		registry:         r,
		blockchainClient: client,
		streamHelper:     h}
}

func (pp *ppDataStreamController) CreateDataStream(ctx context.Context, req resources.CreateDataStreamRequest) resources.CreateDataStreamResponse {
	pp.logger.Printf("entering-create-data-stream")
	defer pp.logger.Printf("exiting-create-data-stream")

	exist, err := pp.blockchainClient.CheckContractID(req.DataContractTypeID, req.Authorization)
	if err != nil {
		pp.logger.Printf("failed-first-attempt-to-check-contract-type-ID-%s-err-%s", req.DataContractTypeID, err.Error())
		time.Sleep(1000 * time.Millisecond)
		exist, err = pp.blockchainClient.CheckContractID(req.DataContractTypeID, req.Authorization)
		if err != nil {
			pp.logger.Printf("failed-second-attempt-to-check-contract-type-ID-%s-err-%s", req.DataContractTypeID, err.Error())
			time.Sleep(1000 * time.Millisecond)
			pp.logger.Println("sleeping a 100 :)")
			exist, err = pp.blockchainClient.CheckContractID(req.DataContractTypeID, req.Authorization)

			if err != nil {
				pp.logger.Printf("failed-thrid-attempt-to-check-contract-type-ID-%s-err-%s", req.DataContractTypeID, err.Error())
				return resources.CreateDataStreamResponse{Error: fmt.Errorf("error-stream-source-failed-to-check-contract-type-id")}
			}
		}
	}
	if !exist {
		pp.logger.Printf("contract-ID-%s-does-not-exist", req.DataContractTypeID)
		return resources.CreateDataStreamResponse{Error: fmt.Errorf("error-contract-id-%s-not-valid", req.DataContractTypeID)}
	}

	if !pp.isValidURL(req.DataStreamSourceURL) {
		pp.logger.Printf("error-stream-source-url-is-not-valid")
		return resources.CreateDataStreamResponse{Error: fmt.Errorf("error-stream-source-url-is-not-valid")}
	}
	if !pp.isValidURL(req.DataStreamDestinationURL) {
		pp.logger.Printf("error-stream-destination-url-is-not-valid")
		return resources.CreateDataStreamResponse{Error: fmt.Errorf("error-stream-destination-url-is-not-valid")}

	}

	dataStream := resources.DataStream{
		DataContractTypeID:       req.DataContractTypeID,
		DataStreamSourceURL:      req.DataStreamSourceURL,
		DataStreamDestinationURL: req.DataStreamDestinationURL,
		DataStreamTopic:          req.DataStreamTopic,
		DataStreamProtocol:       req.DataStreamProtocol,
	}
	err = pp.registry.AddStream(ctx, dataStream)
	if err != nil {
		pp.logger.Printf("error-adding-stream%#v", err)
		return resources.CreateDataStreamResponse{Error: err}
	}

	channel := make(chan string)
	err = pp.registry.AddChannel(ctx, req.DataContractTypeID, channel)
	if err != nil {
		pp.logger.Printf("error-adding-channel%#v", err)
		return resources.CreateDataStreamResponse{Error: err}
	}

	err = pp.streamHelper.StartStreamPullPush(ctx, dataStream, channel)
	if err != nil {
		pp.logger.Printf("error-in-start-stream-%#v", err)
		return resources.CreateDataStreamResponse{Error: err}
	}
	return resources.CreateDataStreamResponse{Stream: dataStream, Error: nil}
}

func (pp *ppDataStreamController) DeleteDataStream(ctx context.Context, req resources.DeleteDataStreamRequest) resources.DeleteDataStreamResponse {
	pp.logger.Printf("entering-delete-data-stream")
	defer pp.logger.Printf("exiting-delete-data-stream")

	exist := pp.registry.HasStream(ctx, req.DataContractTypeID)
	if !exist {
		pp.logger.Printf("error-stream-does-not-exist")
		return resources.DeleteDataStreamResponse{Error: fmt.Errorf("error-stream-does-not-exist")}
	}
	channel, err := pp.registry.GetChannel(ctx, req.DataContractTypeID)
	if err != nil {
		pp.logger.Printf("channel-does-not-exist")
	}

	select {
	case channel <- "close":
		pp.logger.Println("sent close to channel")
	default:
	}
	defer close(channel)

	err = pp.registry.DeleteChannel(ctx, req.DataContractTypeID)
	if err != nil {
		pp.logger.Printf("channel-could-not-be-deleted")
		return resources.DeleteDataStreamResponse{Error: fmt.Errorf("channel-could-not-be-deleted")}
	}

	pp.logger.Printf("deleting-stream")
	err = pp.registry.DeleteStream(ctx, req.DataContractTypeID)
	if err != nil {
		pp.logger.Printf("stream-could-not-be-deleted")
		return resources.DeleteDataStreamResponse{Error: fmt.Errorf("stream-could-not-be-deleted")}
	}

	return resources.DeleteDataStreamResponse{}
}
func (pp *ppDataStreamController) ListDataStreams(ctx context.Context, req resources.ListDataStreamsRequest) resources.ListDataStreamsResponse {
	pp.logger.Printf("entering-list-data-streams")
	defer pp.logger.Printf("exiting-list-data-streams")
	streams := pp.registry.ListStreams()
	return resources.ListDataStreamsResponse{Streams: streams, Error: nil}
}

func (pp *ppDataStreamController) GetDataStream(ctx context.Context, req resources.GetDataStreamRequest) resources.GetDataStreamResponse {
	pp.logger.Printf("entering-get-data-stream")
	defer pp.logger.Printf("exiting-get-data-stream")
	stream, err := pp.registry.GetStream(ctx, req.DataContractTypeID)
	if err != nil {
		pp.logger.Printf("error-getting-stream%#v", err)
		return resources.GetDataStreamResponse{Stream: resources.DataStream{}, Error: err}
	}
	return resources.GetDataStreamResponse{Stream: stream}
}

func (pp *ppDataStreamController) UpdateDataStream(ctx context.Context, req resources.UpdateDataStreamRequest) resources.UpdateDataStreamResponse {
	pp.logger.Printf("entering-update-data-stream")
	defer pp.logger.Printf("exiting-update-data-stream")
	exist := pp.registry.HasStream(ctx, req.DataContractTypeID)
	if !exist {
		pp.logger.Printf("stream-does-not-exist")
		return resources.UpdateDataStreamResponse{Stream: resources.DataStream{}, Error: fmt.Errorf("stream-does-not-exist")}
	}
	dataStream := resources.DataStream{
		DataContractTypeID:       req.DataContractTypeID,
		DataStreamSourceURL:      req.DataStreamSourceURL,
		DataStreamDestinationURL: req.DataStreamDestinationURL,
		DataStreamTopic:          req.DataStreamTopic,
		DataStreamProtocol:       req.DataStreamProtocol,
	}

	err := pp.registry.UpdateStream(ctx, dataStream)
	if err != nil {
		pp.logger.Printf("error-updating-registry%#v", err)
		return resources.UpdateDataStreamResponse{Error: err}
	}

	channel, err := pp.registry.GetChannel(ctx, req.DataContractTypeID)
	if err != nil {
		pp.logger.Printf("error-getting-channel-from-registry%#v", err)
		return resources.UpdateDataStreamResponse{Error: err}
	}

	select {
	case channel <- "close":
		pp.logger.Println("sent close to channel")
	default:
	}

	err = pp.streamHelper.StartStreamPullPush(ctx, dataStream, channel)
	if err != nil {
		pp.logger.Printf("error-in-restart-stream-%#v", err)
		return resources.UpdateDataStreamResponse{Error: err}
	}
	return resources.UpdateDataStreamResponse{Stream: dataStream}
}

//isValidURL
func (pp *ppDataStreamController) isValidURL(stringToTest string) bool {
	_, err := url.ParseRequestURI(stringToTest)
	if err != nil {
		return false
	}
	return true

}
