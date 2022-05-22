package main

import (
	"fmt"
	"os"

	"github.com/nndergunov/messageBoard/server/api"
	"github.com/nndergunov/messageBoard/server/api/v1/handlers"
	"github.com/nndergunov/messageBoard/server/cmd/server"
	"github.com/nndergunov/messageBoard/server/cmd/server/config"
	"github.com/nndergunov/messageBoard/server/pkg/app"
	"github.com/nndergunov/messageBoard/server/pkg/configreader"
	"github.com/nndergunov/messageBoard/server/pkg/db"
	"github.com/nndergunov/messageBoard/server/pkg/logger"
)

func main() {
	mainLogger := logger.NewLogger(os.Stdout, "main")

	err := configreader.SetConfigFile("config.yaml")
	if err != nil {
		mainLogger.Panicln(err)
	}

	dbURL := fmt.Sprintf(
		"host=" + configreader.GetString("database.host") +
			" port=" + configreader.GetString("database.port") +
			" user=" + configreader.GetString("database.user") +
			" password=" + configreader.GetString("database.password") +
			" dbname=" + configreader.GetString("database.dbname") +
			" sslmode=" + configreader.GetString("database.ssl"),
	)

	database, err := db.NewDB(dbURL)
	if err != nil {
		mainLogger.Panicln(err)
	}

	appInstance := app.NewApp(database)

	handlerLogger := logger.NewLogger(os.Stdout, "handlers")
	handler := handlers.NewEndpointHandler(appInstance, handlerLogger)

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

	messageBrokerServer := server.NewServer(serverConfig)

	stopChan := make(chan interface{})

	messageBrokerServer.StartListening(stopChan)

	<-stopChan
}
