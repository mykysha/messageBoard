package handlers

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/messageBoard/server/api/v1"
	"github.com/nndergunov/messageBoard/server/pkg/logger"
)

type endpointHandler struct {
	router *mux.Router
	log    *logger.Logger
}

// NewEndpointHandler returns new http multiplexer with configured endpoints.
func NewEndpointHandler(log *logger.Logger) *mux.Router {
	router := mux.NewRouter()

	handler := endpointHandler{
		router: router,
		log:    log,
	}

	handler.handlerInit()

	return handler.router
}

func (e *endpointHandler) handlerInit() {
	e.router.HandleFunc("/v1/status", e.statusHandler)
}

func (e endpointHandler) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "accounting",
		IsUp:        "up",
	}

	status, err := v1.EncodeIndent(data, "", " ")
	if err != nil {
		e.log.Println(err)
	}

	_, err = io.WriteString(responseWriter, string(status))
	if err != nil {
		e.log.Printf("\nstatus write: %v", err)

		return
	}

	e.log.Printf("\ngave status %s", data.IsUp)
}
