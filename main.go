//
// Copyright (c) 2019 LG Electronics Inc.
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"context"

	"github.com/lgsvl/data-marketplace-stream-delivery/controllers"
	"github.com/lgsvl/data-marketplace-stream-delivery/resources"
	"github.com/lgsvl/data-marketplace-stream-delivery/server"
	"github.com/lgsvl/data-marketplace-stream-delivery/utils"
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
