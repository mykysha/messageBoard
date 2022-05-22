package main

import (
	"os"

	"github.com/nndergunov/messageBoard/server/api"
	"github.com/nndergunov/messageBoard/server/api/v1/handlers"
	"github.com/nndergunov/messageBoard/server/cmd/server"
	"github.com/nndergunov/messageBoard/server/cmd/server/config"
	"github.com/nndergunov/messageBoard/server/pkg/configreader"
	"github.com/nndergunov/messageBoard/server/pkg/logger"
)

func main() {
	mainLogger := logger.NewLogger(os.Stdout, "main")

	err := configreader.SetConfigFile("config.yaml")
	if err != nil {
		mainLogger.Panicln(err)
	}

	handlerLogger := logger.NewLogger(os.Stdout, "handlers")
	handler := handlers.NewEndpointHandler(handlerLogger)

	apiLogger := logger.NewLogger(os.Stdout, "handlers")
	serverAPI := api.NewAPI(handler, apiLogger)

	serverLogger := logger.NewLogger(os.Stdout, "server")
	serverConfig := &config.Config{
		Address:           configreader.GetString("server-config.address"),
		ReadTimeout:       configreader.GetDuration("server-config.readTime"),
		WriteTimeout:      configreader.GetDuration("server-config.writeTime"),
		IdleTimeout:       configreader.GetDuration("server-config.idleTime"),
		ReadHeaderTimeout: configreader.GetDuration("server-config.readerHeaderTime"),
		ErrorLog:          nil,
		ServerLogger:      serverLogger,
		API:               serverAPI,
	}

	proxyServer := server.NewServer(serverConfig)

	stopChan := make(chan interface{})

	proxyServer.StartListening(stopChan)

	<-stopChan
}
