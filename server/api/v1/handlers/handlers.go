package handlers

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/messageBoard/server/api/v1"
	"github.com/nndergunov/messageBoard/server/pkg/app"
	"github.com/nndergunov/messageBoard/server/pkg/logger"
)

type endpointHandler struct {
	router      *mux.Router
	appInstance *app.App
	log         *logger.Logger
}

// NewEndpointHandler returns new http multiplexer with configured endpoints.
func NewEndpointHandler(appInstance *app.App, log *logger.Logger) *mux.Router {
	router := mux.NewRouter()

	handler := endpointHandler{
		router:      router,
		appInstance: appInstance,
		log:         log,
	}

	handler.handlerInit()

	return handler.router
}

func (e *endpointHandler) handlerInit() {
	e.router.HandleFunc("/v1/status", e.statusHandler)
	e.router.HandleFunc("/v1/messages", e.returnAllMessages).Methods(http.MethodGet)
	e.router.HandleFunc("/v1/messages", e.createMessage).Methods(http.MethodPost)
}

func (e endpointHandler) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		IsUp: "up",
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

func (e endpointHandler) returnAllMessages(responseWriter http.ResponseWriter, _ *http.Request) {
	messages, err := e.appInstance.ReturnAllMessages()
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	e.respond(messages, responseWriter)
}

func (e endpointHandler) createMessage(responseWriter http.ResponseWriter, request *http.Request) {
	message, err := requestToMessage(request)
	if err != nil {
		if err != nil {
			e.log.Println(err)

			responseWriter.WriteHeader(http.StatusBadRequest)

			return
		}
	}

	err = e.appInstance.InsertMessage(*message)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}
