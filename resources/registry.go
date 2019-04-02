package resources

import (
	"context"
	"fmt"
	"log"
)

//Registry is a registry that saves the created data streams
//go:generate counterfeiter -o ../fakes/fake_registry.go . Registry
type Registry interface {
	AddStream(context.Context, DataStream) error
	UpdateStream(context.Context, DataStream) error
	HasStream(context.Context, string) bool
	GetStream(context.Context, string) (DataStream, error)
	DeleteStream(context.Context, string) error
	ListStreams() []DataStream
	AddChannel(context.Context, string, chan string) error
	GetChannel(context.Context, string) (chan string, error)
	DeleteChannel(context.Context, string) error
}

type registry struct {
	ctx         context.Context
	logger      *log.Logger
	streamsMap  map[string]DataStream
	channelsMap map[string]chan string
}

//NewRegistry creates a new empty registry
func NewRegistry(ctx context.Context, registryLogger *log.Logger) Registry {
	return &registry{
		logger:      registryLogger,
		streamsMap:  make(map[string]DataStream),
		channelsMap: make(map[string]chan string),
	}
}

//AddStream adds a new data stream to the registry
func (r *registry) AddStream(ctx context.Context, stream DataStream) error {
	r.logger.Printf("entering-registry-add-stream")
	defer r.logger.Printf("exiting-registry-add-stream")
	_, ok := r.streamsMap[stream.DataContractTypeID]
	if ok {
		r.logger.Printf("error-adding-stream-uuid-exists")
		return fmt.Errorf("error-adding-stream-uuid-exists")
	}
	r.streamsMap[stream.DataContractTypeID] = stream
	return nil
}

//UpdateStream updates a data stream in the registry
func (r *registry) UpdateStream(ctx context.Context, stream DataStream) error {
	r.logger.Printf("entering-registry-update-stream")
	defer r.logger.Printf("exiting-registry-update-stream")
	_, ok := r.streamsMap[stream.DataContractTypeID]
	if !ok {
		r.logger.Printf("error-updating-stream-does-not-exist")
		return fmt.Errorf("error-updating-stream-does-not-exist")
	}
	r.streamsMap[stream.DataContractTypeID] = stream
	return nil
}

//HasElement returns if the registry contains a specific element
func (r *registry) HasStream(ctx context.Context, dataStreamContractID string) bool {
	r.logger.Printf("entering-registry-has-stream")
	defer r.logger.Printf("exiting-registry-has-stream")
	_, ok := r.streamsMap[dataStreamContractID]
	return ok
}

//GetStream returns a stream instance if the registry has it
func (r *registry) GetStream(ctx context.Context, dataStreamContractID string) (DataStream, error) {
	r.logger.Printf("entering-registry-get-stream")
	defer r.logger.Printf("exiting-registry-get-stream")
	stream, ok := r.streamsMap[dataStreamContractID]
	if !ok {
		r.logger.Printf("error-getting-stream-stream-does-not-exist")
		return DataStream{}, fmt.Errorf("error-getting-stream-stream-does-not-exist")
	}
	return stream, nil
}

func (r *registry) ListStreams() []DataStream {
	r.logger.Printf("entering-registry-list-streams")
	defer r.logger.Printf("exiting-registry-list-streams")
	resp := []DataStream{}
	for _, str := range r.streamsMap {
		resp = append(resp, str)
	}

	return resp
}

//DeleteStream deletes a data stream from the registry
func (r *registry) DeleteStream(ctx context.Context, streamID string) error {
	r.logger.Printf("entering-registry-delete-stream")
	defer r.logger.Printf("exiting-registry-delete-stream")
	_, ok := r.streamsMap[streamID]
	if !ok {
		r.logger.Printf("error-deleting-stream-does-not-exist")
		return fmt.Errorf("error-deleting-stream-does-not-exist")
	}
	delete(r.streamsMap, streamID)
	return nil
}

//AddChannel adds a new channel  to the registry
func (r *registry) AddChannel(ctx context.Context, streamUUID string, channel chan string) error {
	r.logger.Printf("entering-registry-add-channel")
	defer r.logger.Printf("exiting-registry-add-channel")

	_, ok := r.streamsMap[streamUUID]
	if !ok {
		r.logger.Printf("error-adding-channel-uuid-does-not-exist-in-streams")
		return fmt.Errorf("error-adding-channel-uuid-does-not-exist-in-streams")
	}
	_, ok = r.channelsMap[streamUUID]
	if ok {
		r.logger.Printf("error-adding-channel-uuid-exists")
		return fmt.Errorf("error-adding-channel-uuid-exists")
	}
	r.channelsMap[streamUUID] = channel
	return nil
}

//GetChannel gets a channel from the registry
func (r *registry) GetChannel(ctx context.Context, streamUUID string) (chan string, error) {
	r.logger.Printf("entering-registry-get-channel")
	defer r.logger.Printf("exiting-registry-get-channel")
	channel, ok := r.channelsMap[streamUUID]
	if !ok {
		r.logger.Printf("error-getting-channel-uuid-does-not-exists")
		return nil, fmt.Errorf("error-getting-channel-uuid-does-not-exists")
	}

	return channel, nil
}

//DeleteChannel deletes a channel from the registry
func (r *registry) DeleteChannel(ctx context.Context, streamUUID string) error {
	r.logger.Printf("entering-registry-delete-channel")
	defer r.logger.Printf("exiting-registry-delete-channel")
	_, ok := r.channelsMap[streamUUID]
	if !ok {
		r.logger.Printf("error-deleting-channel-uuid-does-not-exists")
		return fmt.Errorf("error-deleting-channel-uuid-does-not-exists")
	}
	delete(r.channelsMap, streamUUID)
	return nil
}
