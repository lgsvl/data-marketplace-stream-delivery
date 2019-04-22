//
// Copyright (c) 2019 LG Electronics Inc.
// SPDX-License-Identifier: Apache-2.0
//

package resources

import (
	"context"
)

const (
	//CheckContractIDPath has the path tp check that the contractID exists in the blockchain
	CheckContractIDPath string = "api/DataContractType/"
)

//DataStreamController controls actions on DataStreams
//go:generate counterfeiter -o ../fakes/fake_data_stream_controller.go . DataStreamController
type DataStreamController interface {
	CreateDataStream(context.Context, CreateDataStreamRequest) CreateDataStreamResponse
	DeleteDataStream(context.Context, DeleteDataStreamRequest) DeleteDataStreamResponse
	ListDataStreams(context.Context, ListDataStreamsRequest) ListDataStreamsResponse
	GetDataStream(context.Context, GetDataStreamRequest) GetDataStreamResponse
	UpdateDataStream(context.Context, UpdateDataStreamRequest) UpdateDataStreamResponse
}

//ServerConfig contains the configuration of the server
type ServerConfig struct {
	Port int
}

//CreateDataStreamRequest contains the data stream request parameters
type CreateDataStreamRequest struct {
	Authorization            map[string]string `json:"authorization,omitempty"`
	DataContractTypeID       string            `json:"DataContractTypeID,omitempty"`
	DataStreamProtocol       string            `json:"DataStreamProtocol,omitempty"`
	DataStreamSourceURL      string            `json:"DataStreamSourceURL,omitempty"`
	DataStreamDestinationURL string            `json:"DataStreamDestinationURL,omitempty"`
	DataStreamTopic          string            `json:"DataStreamTopic,omitempty"`
	// Duration StreamDuration = 5;
	DataStreamParameters map[string]string `json:"DataStreamParameters,omitempty"`
}

//CreateDataStreamResponse contains a response to create data stream
type CreateDataStreamResponse struct {
	Stream DataStream
	Error  error
}

//DataStream contains the data stream request parameters
type DataStream struct {
	DataContractTypeID       string `json:"DataContractTypeID,omitempty"`
	DataStreamProtocol       string `json:"DataStreamProtocol,omitempty"`
	DataStreamSourceURL      string `json:"DataStreamSourceURL,omitempty"`
	DataStreamDestinationURL string `json:"DataStreamDestinationURL,omitempty"`
	DataStreamTopic          string `json:"DataStreamTopic,omitempty"`
	// Duration StreamDuration = 5;
	DataStreamParameters map[string]string `json:"DataStreamParameters,omitempty"`
}

//DeleteDataStreamRequest contains a request to delete a stream
type DeleteDataStreamRequest struct {
	DataContractTypeID string `json:"DataContractTypeID,omitempty"`
}

//DeleteDataStreamResponse a response to delete data stream with any error
type DeleteDataStreamResponse struct {
	Error error `json:"error"`
}

//ListDataStreamsRequest contains a request to list all streams
type ListDataStreamsRequest struct {
	Filter string `json:"filter,omitempty"`
}

//ListDataStreamsResponse a response to list data streams with an array of existing streams
type ListDataStreamsResponse struct {
	Streams []DataStream `json:"streams"`
	Error   error        `json:"error"`
}

//GetDataStreamRequest contains a request to get a stream
type GetDataStreamRequest struct {
	DataContractTypeID string `json:"DataContractTypeID"`
}

//GetDataStreamResponse a response to get data stream with a stream
type GetDataStreamResponse struct {
	Stream DataStream `json:"stream"`
	Error  error      `json:"error"`
}

//UpdateDataStreamRequest contains a request to update a stream
type UpdateDataStreamRequest struct {
	DataContractTypeID       string `json:"DataContractTypeID,omitempty"`
	DataStreamProtocol       string `json:"DataStreamProtocol,omitempty"`
	DataStreamSourceURL      string `json:"DataStreamSourceURL,omitempty"`
	DataStreamDestinationURL string `json:"DataStreamDestinationURL,omitempty"`
	DataStreamTopic          string `json:"DataStreamTopic,omitempty"`
	// Duration StreamDuration = 5;
	DataStreamParameters map[string]string `json:"DataStreamParameters,omitempty"`
}

//UpdateDataStreamResponse a response to update data stream
type UpdateDataStreamResponse struct {
	Stream DataStream `json:"stream"`
	Error  error      `json:"error"`
}
