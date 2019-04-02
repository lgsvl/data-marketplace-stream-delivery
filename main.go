package main

import (
	"context"

	"pse-gitlab.lgsvl.net/data_marketplace/data-stream-delivery/controllers"
	"pse-gitlab.lgsvl.net/data_marketplace/data-stream-delivery/resources"
	"pse-gitlab.lgsvl.net/data_marketplace/data-stream-delivery/server"
	"pse-gitlab.lgsvl.net/data_marketplace/data-stream-delivery/utils"
)

// main start a http server and waits for connection
func main() {
	ctx := context.Background()
	config := resources.ServerConfig{Port: 7778}
	logger := utils.CreateLogger("data-stream-delivery")
	registry := resources.NewRegistry(ctx, logger)
	ppController := controllers.NewPPDataStreamController(ctx, logger, registry)
	handler := server.NewDataStreamHandler(ctx, logger, ppController)
	server := server.NewDataStreamServer(ctx, logger, handler, config)
	logger.Printf("starting-server", server.Start())
}
