//
// Copyright (c) 2019 LG Electronics Inc.
// SPDX-License-Identifier: Apache-2.0
//

package controllers_test

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/lgsvl/data-marketplace-stream-delivery/controllers"
	"github.com/lgsvl/data-marketplace-stream-delivery/fakes"
	"github.com/lgsvl/data-marketplace-stream-delivery/resources"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("pull-push-controller", func() {
	var (
		logger           *log.Logger
		fakeRegistry     *fakes.FakeRegistry
		fakeClient       *fakes.FakeBlockchainClient
		fakeStreamHelper *fakes.FakeStreamHelper
		ctx              context.Context
		controller       resources.DataStreamController
	)
	BeforeEach(func() {
		logger = log.New(os.Stdout, "data-stream-delivery-controller-test: ", log.Lshortfile|log.LstdFlags)
		ctx = context.Background()
		fakeRegistry = new(fakes.FakeRegistry)
		fakeClient = new(fakes.FakeBlockchainClient)
		fakeStreamHelper = new(fakes.FakeStreamHelper)
		controller = controllers.NewPPDataStreamControllerWithClients(ctx, logger, fakeRegistry, fakeClient, fakeStreamHelper)

	})

	Context(".CreateDataStream", func() {
		It("should fail when blockchain client fails to check contractID", func() {
			request := resources.CreateDataStreamRequest{DataContractTypeID: "123",
				DataStreamSourceURL:      "fakeSRC",
				DataStreamDestinationURL: "fakeDest",
				DataStreamProtocol:       "fakeProtocl",
				DataStreamTopic:          "fakeTopic"}
			fakeClient.CheckContractIDReturns(false, fmt.Errorf("error-checking"))
			response := controller.CreateDataStream(ctx, request)

			Expect(response.Error).To(HaveOccurred())
			Expect(fakeRegistry.AddStreamCallCount()).To(Equal(0))

		})
		It("should fail when blockchain client does not find contractID", func() {
			request := resources.CreateDataStreamRequest{DataContractTypeID: "123",
				DataStreamSourceURL:      "fakeSRC",
				DataStreamDestinationURL: "fakeDest",
				DataStreamProtocol:       "fakeProtocl",
				DataStreamTopic:          "fakeTopic"}
			fakeClient.CheckContractIDReturns(false, nil)
			response := controller.CreateDataStream(ctx, request)

			Expect(response.Error).To(HaveOccurred())
			Expect(fakeRegistry.AddStreamCallCount()).To(Equal(0))

		})

		It("should fail when source url is not valid", func() {
			request := resources.CreateDataStreamRequest{DataContractTypeID: "123",
				DataStreamSourceURL:      "fakeSRC",
				DataStreamDestinationURL: "fakeDest",
				DataStreamProtocol:       "fakeProtocl",
				DataStreamTopic:          "fakeTopic"}
			fakeClient.CheckContractIDReturns(true, nil)

			response := controller.CreateDataStream(ctx, request)

			Expect(response.Error).To(HaveOccurred())
			Expect(fakeRegistry.AddStreamCallCount()).To(Equal(0))

		})
		It("should fail when destination url is not valid", func() {
			request := resources.CreateDataStreamRequest{DataContractTypeID: "123",
				DataStreamSourceURL:      "fakeSRC",
				DataStreamDestinationURL: "fakeDest",
				DataStreamProtocol:       "fakeProtocl",
				DataStreamTopic:          "fakeTopic"}

			fakeClient.CheckContractIDReturns(true, nil)
			fakeRegistry.AddStreamReturns(fmt.Errorf("stream-already-exists"))
			response := controller.CreateDataStream(ctx, request)

			Expect(response.Error).To(HaveOccurred())
			Expect(fakeRegistry.AddStreamCallCount()).To(Equal(0))

		})

		It("should fail when registry fails to add stream", func() {
			request := resources.CreateDataStreamRequest{DataContractTypeID: "123",
				DataStreamSourceURL:      "http://test.com",
				DataStreamDestinationURL: "destination:9090",
				DataStreamProtocol:       "fakeProtocl",
				DataStreamTopic:          "fakeTopic"}
			fakeRegistry.AddStreamReturns(fmt.Errorf("error-adding-stream"))

			fakeClient.CheckContractIDReturns(true, nil)
			response := controller.CreateDataStream(ctx, request)

			Expect(response.Error).To(HaveOccurred())
			Expect(fakeRegistry.AddStreamCallCount()).To(Equal(1))
			Expect(fakeRegistry.AddChannelCallCount()).To(Equal(0))

		})

		It("should fail when registry succeeds to add stream and fails to add channel", func() {
			request := resources.CreateDataStreamRequest{DataContractTypeID: "123",
				DataStreamSourceURL:      "http://test.com",
				DataStreamDestinationURL: "destination:9090",
				DataStreamProtocol:       "fakeProtocl",
				DataStreamTopic:          "fakeTopic"}

			fakeClient.CheckContractIDReturns(true, nil)
			fakeRegistry.AddStreamReturns(nil)
			fakeRegistry.AddChannelReturns(fmt.Errorf("error-adding-channel"))
			response := controller.CreateDataStream(ctx, request)

			Expect(response.Error).To(HaveOccurred())
			Expect(fakeRegistry.AddStreamCallCount()).To(Equal(1))
			Expect(fakeRegistry.AddChannelCallCount()).To(Equal(1))

		})

		It("should fail when stream helper fails to start stream", func() {
			request := resources.CreateDataStreamRequest{DataContractTypeID: "123",
				DataStreamSourceURL:      "http://test.com",
				DataStreamDestinationURL: "fakedestination:9090",
				DataStreamProtocol:       "fakeProtocl",
				DataStreamTopic:          "fakeTopic"}

			fakeClient.CheckContractIDReturns(true, nil)
			fakeRegistry.AddStreamReturns(nil)
			fakeRegistry.AddChannelReturns(nil)
			fakeStreamHelper.StartStreamPullPushReturns(fmt.Errorf("error-starting-stream"))
			response := controller.CreateDataStream(ctx, request)

			Expect(response.Error).To(HaveOccurred())
			Expect(fakeRegistry.AddStreamCallCount()).To(Equal(1))
			Expect(fakeRegistry.AddChannelCallCount()).To(Equal(1))
			Expect(fakeStreamHelper.StartStreamPullPushCallCount()).To(Equal(1))

		})

		It("should succeed when stream helper succeeds to start stream", func() {
			request := resources.CreateDataStreamRequest{DataContractTypeID: "123",
				DataStreamSourceURL:      "http://test.com",
				DataStreamDestinationURL: "fakedestination:9090",
				DataStreamProtocol:       "fakeProtocl",
				DataStreamTopic:          "fakeTopic"}

			fakeClient.CheckContractIDReturns(true, nil)
			fakeRegistry.AddStreamReturns(nil)
			fakeRegistry.AddChannelReturns(nil)
			fakeStreamHelper.StartStreamPullPushReturns(nil)
			response := controller.CreateDataStream(ctx, request)

			Expect(response.Error).NotTo(HaveOccurred())
			Expect(fakeRegistry.AddStreamCallCount()).To(Equal(1))
			Expect(fakeRegistry.AddChannelCallCount()).To(Equal(1))
			Expect(fakeStreamHelper.StartStreamPullPushCallCount()).To(Equal(1))

		})
	})

	Context(".DeleteDataStream", func() {
		It("should fail when data stream contractID does not exist", func() {
			request := resources.DeleteDataStreamRequest{DataContractTypeID: "123"}
			fakeRegistry.HasStreamReturns(false)
			response := controller.DeleteDataStream(ctx, request)

			Expect(response.Error).To(HaveOccurred())
			Expect(fakeRegistry.DeleteStreamCallCount()).To(Equal(0))

		})
		It("should fail when registry cannot delete channel", func() {
			request := resources.DeleteDataStreamRequest{DataContractTypeID: "123"}
			fakeRegistry.HasStreamReturns(true)
			fakeRegistry.GetChannelReturns(make(chan string), nil)
			fakeRegistry.DeleteChannelReturns(fmt.Errorf("failed-deleting-channel"))
			response := controller.DeleteDataStream(ctx, request)

			Expect(response.Error).To(HaveOccurred())
			Expect(fakeRegistry.DeleteChannelCallCount()).To(Equal(1))
			Expect(fakeRegistry.DeleteStreamCallCount()).To(Equal(0))

		})
		It("should fail when registry cannot delete stream", func() {
			request := resources.DeleteDataStreamRequest{DataContractTypeID: "123"}
			fakeRegistry.HasStreamReturns(true)
			fakeRegistry.GetChannelReturns(make(chan string), nil)
			fakeRegistry.DeleteChannelReturns(nil)
			fakeRegistry.DeleteStreamReturns(fmt.Errorf("failed-deleting-channel"))
			response := controller.DeleteDataStream(ctx, request)

			Expect(response.Error).To(HaveOccurred())
			Expect(fakeRegistry.DeleteChannelCallCount()).To(Equal(1))
			Expect(fakeRegistry.DeleteStreamCallCount()).To(Equal(1))

		})

		It("should succeed when registry can delete stream", func() {
			request := resources.DeleteDataStreamRequest{DataContractTypeID: "123"}
			fakeRegistry.HasStreamReturns(true)
			fakeRegistry.GetChannelReturns(make(chan string), nil)
			fakeRegistry.DeleteChannelReturns(nil)
			fakeRegistry.DeleteStreamReturns(nil)
			response := controller.DeleteDataStream(ctx, request)

			Expect(response.Error).NotTo(HaveOccurred())
			Expect(fakeRegistry.DeleteChannelCallCount()).To(Equal(1))
			Expect(fakeRegistry.DeleteStreamCallCount()).To(Equal(1))

		})
	})

	Context(".ListDataStreams", func() {
		It("should succedd", func() {
			streams := []resources.DataStream{}
			streams = append(streams, resources.DataStream{})
			fakeRegistry.ListStreamsReturns(streams)
			req := resources.ListDataStreamsRequest{}
			response := controller.ListDataStreams(ctx, req)
			Expect(response.Error).NotTo(HaveOccurred())
			Expect(len(response.Streams)).To(Equal(1))

		})
	})

	Context(".GetDataStream", func() {
		It("should fail when registry fails to get stream", func() {

			fakeRegistry.GetStreamReturns(resources.DataStream{}, fmt.Errorf("error"))
			req := resources.GetDataStreamRequest{}
			response := controller.GetDataStream(ctx, req)
			Expect(response.Error).To(HaveOccurred())

		})
		It("should succeed when registry succeeds to get stream", func() {

			fakeRegistry.GetStreamReturns(resources.DataStream{}, nil)
			req := resources.GetDataStreamRequest{}
			response := controller.GetDataStream(ctx, req)
			Expect(response.Error).NotTo(HaveOccurred())

		})
	})

	Context(".UpdateDataStream", func() {
		It("should fail when registry does not have stream", func() {

			fakeRegistry.HasStreamReturns(false)
			req := resources.UpdateDataStreamRequest{DataContractTypeID: "123"}
			response := controller.UpdateDataStream(ctx, req)
			Expect(response.Error).To(HaveOccurred())
			Expect(fakeRegistry.UpdateStreamCallCount()).To(Equal(0))

		})

		It("should fail when registry fails to update stream", func() {
			fakeRegistry.HasStreamReturns(true)
			fakeRegistry.UpdateStreamReturns(fmt.Errorf("error"))
			req := resources.UpdateDataStreamRequest{DataContractTypeID: "123"}
			response := controller.UpdateDataStream(ctx, req)
			Expect(response.Error).To(HaveOccurred())
			Expect(fakeRegistry.UpdateStreamCallCount()).To(Equal(1))

		})
		It("should succeed when registry succeeds to update stream", func() {
			fakeRegistry.HasStreamReturns(true)
			fakeRegistry.UpdateStreamReturns(nil)
			req := resources.UpdateDataStreamRequest{DataContractTypeID: "123"}
			response := controller.UpdateDataStream(ctx, req)
			Expect(response.Error).NotTo(HaveOccurred())
			Expect(response.Stream.DataContractTypeID).To(Equal("123"))

		})
	})
})
