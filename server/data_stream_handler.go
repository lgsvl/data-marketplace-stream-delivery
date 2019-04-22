//
// Copyright (c) 2019 LG Electronics Inc.
// SPDX-License-Identifier: Apache-2.0
//

package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/lgsvl/data-marketplace-stream-delivery/resources"
	"github.com/lgsvl/data-marketplace-stream-delivery/utils"
)

//DataStreamHandler is a handler for data streams
type DataStreamHandler struct {
	ctx        context.Context
	logger     *log.Logger
	controller resources.DataStreamController
}

//NewDataStreamHandler returns a new DataStreamHandler
func NewDataStreamHandler(c context.Context, lg *log.Logger, cont resources.DataStreamController) *DataStreamHandler {
	return &DataStreamHandler{ctx: c, logger: lg, controller: cont}
}

//CreateDataStream creates a handlerFunc for crating a new PullPush DataStream
func (d *DataStreamHandler) CreateDataStream() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		d.logger.Printf("entering-create-data-stream")
		defer d.logger.Printf("exiting-create-push-pull-data-stream")
		d.EnableCORS(w, req)

		ctx := context.Background()
		createDataStreamRequest := resources.CreateDataStreamRequest{}
		err := utils.UnmarshalDataFromRequest(req, &createDataStreamRequest)
		if err != nil {
			utils.WriteResponse(w, 409, &resources.CreateDataStreamResponse{Error: err})
			return
		}
		auth := utils.GetFromHeader(req, "authorization")
		createDataStreamRequest.Authorization = map[string]string{"authorization": auth}

		resp := d.controller.CreateDataStream(ctx, createDataStreamRequest)
		if resp.Error != nil {
			utils.WriteResponse(w, 409, &resp)
			return
		}
		utils.WriteResponse(w, http.StatusOK, resp)
	}
}

//DeleteDataStream creates a handlerFunc for deleting a PullPush DataStream
func (d *DataStreamHandler) DeleteDataStream() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		d.logger.Printf("entering-delete-data-stream")
		defer d.logger.Printf("exiting-delete-push-pull-data-stream")
		d.EnableCORS(w, req)

		ctx := context.Background()
		streamID := utils.ExtractVarsFromRequest(req, "streamID")
		deleteDataStreamRequest := resources.DeleteDataStreamRequest{DataContractTypeID: streamID}

		resp := d.controller.DeleteDataStream(ctx, deleteDataStreamRequest)
		if resp.Error != nil {
			utils.WriteResponse(w, 409, &resp)
			return
		}
		utils.WriteResponse(w, http.StatusOK, resp)
	}
}

//ListDataStreams creates a handlerFunc for listing all PullPush DataStreams
func (d *DataStreamHandler) ListDataStreams() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		d.logger.Printf("entering-list-data-streams")
		defer d.logger.Printf("exiting-list-data-streams")
		d.EnableCORS(w, req)

		ctx := context.Background()
		listDataStreamsRequest := resources.ListDataStreamsRequest{}
		resp := d.controller.ListDataStreams(ctx, listDataStreamsRequest)
		if resp.Error != nil {
			utils.WriteResponse(w, 409, &resp)
			return
		}
		utils.WriteResponse(w, http.StatusOK, &resp)
	}
}

//GetDataStream creates a handlerFunc for getting a DataStream
func (d *DataStreamHandler) GetDataStream() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		d.logger.Printf("entering-get-data-stream")
		defer d.logger.Printf("exiting-get-data-stream")
		d.EnableCORS(w, req)

		ctx := context.Background()
		streamID := utils.ExtractVarsFromRequest(req, "streamID")
		getDataStreamRequest := resources.GetDataStreamRequest{DataContractTypeID: streamID}

		resp := d.controller.GetDataStream(ctx, getDataStreamRequest)
		if resp.Error != nil {
			utils.WriteResponse(w, 409, &resp)
			return
		}
		utils.WriteResponse(w, http.StatusOK, &resp)
	}
}

//UpdateDataStream creates a handlerFunc for getting a DataStream
func (d *DataStreamHandler) UpdateDataStream() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		d.logger.Printf("entering-update-data-stream")
		defer d.logger.Printf("exiting-update-data-stream")
		d.EnableCORS(w, req)

		ctx := context.Background()
		streamID := utils.ExtractVarsFromRequest(req, "streamID")
		updateDataStreamRequest := resources.UpdateDataStreamRequest{}
		err := utils.UnmarshalDataFromRequest(req, &updateDataStreamRequest)
		if err != nil {
			utils.WriteResponse(w, 409, &resources.UpdateDataStreamResponse{Error: err})
			return
		}

		if streamID != updateDataStreamRequest.DataContractTypeID {
			utils.WriteResponse(w, 409, &resources.UpdateDataStreamResponse{Error: fmt.Errorf("streamID url parameter should match body DataContractTypeID")})
			return
		}
		resp := d.controller.UpdateDataStream(ctx, updateDataStreamRequest)
		if resp.Error != nil {
			utils.WriteResponse(w, 409, &resp)
			return
		}
		utils.WriteResponse(w, http.StatusOK, &resp)
	}
}

//OptionsProcessing Handles preflight checks from portal
func (d *DataStreamHandler) OptionsProcessing() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		d.logger.Printf("entering-options-processing")
		defer d.logger.Printf("exiting-options-processing")
		d.EnableCORS(w, req)
		d.logger.Printf("received-options")
		utils.WriteResponse(w, http.StatusOK, "")
		return
	}
}

//EnableCORS to allow Origin
func (d *DataStreamHandler) EnableCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	return
}
