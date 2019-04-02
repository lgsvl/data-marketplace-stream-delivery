package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lgsvl/data-marketplace-stream-delivery/resources"
)

// DataStreamServer represents the DataStream server
type DataStreamServer struct {
	ctx               context.Context
	logger            *log.Logger
	dataStreamHandler *DataStreamHandler
	config            resources.ServerConfig
}

//NewDataStreamServer creates a new instance of the DataStreamServer
func NewDataStreamServer(c context.Context, l *log.Logger, d *DataStreamHandler, conf resources.ServerConfig) *DataStreamServer {
	return &DataStreamServer{ctx: c, logger: l, dataStreamHandler: d, config: conf}
}

//InitializeHandler prepares the handlers to be used for the rest queries
func (d *DataStreamServer) InitializeHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/v1/pullpushstream", d.dataStreamHandler.CreateDataStream()).Methods("POST")
	router.HandleFunc("/v1/pullpushstream", d.dataStreamHandler.OptionsProcessing()).Methods("OPTIONS")
	router.HandleFunc("/v1/pullpushstream/{streamID}", d.dataStreamHandler.DeleteDataStream()).Methods("DELETE")
	router.HandleFunc("/v1/pullpushstream/{streamID}", d.dataStreamHandler.GetDataStream()).Methods("GET")
	router.HandleFunc("/v1/pullpushstream/{streamID}", d.dataStreamHandler.UpdateDataStream()).Methods("PUT")

	router.HandleFunc("/v1/pullpushstreams", d.dataStreamHandler.ListDataStreams()).Methods("GET")

	return router
}

//Start starts the http server
func (d *DataStreamServer) Start() error {
	router := d.InitializeHandler()
	http.Handle("/", router)
	d.printStartMsg()

	return http.ListenAndServe(fmt.Sprintf(":%d", d.config.Port), nil)
}
func (d *DataStreamServer) printStartMsg() {
	d.logger.Printf(fmt.Sprintf("Starting data stream server on port %d ....", d.config.Port))
	d.logger.Printf("CTL-C to exit/stop data-stream-delivery service")
}
