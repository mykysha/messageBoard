package handlers

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/goproxy/api/v1"
	"github.com/nndergunov/goproxy/pkg/app/proxyapp"
	"github.com/nndergunov/goproxy/pkg/logger"
)

type endpointHandler struct {
	proxy  *proxyapp.ProxyApp
	router *mux.Router
	log    *logger.Logger
}

// NewEndpointHandler returns new http multiplexer with configured endpoints.
func NewEndpointHandler(proxy *proxyapp.ProxyApp, log *logger.Logger) *mux.Router {
	router := mux.NewRouter()

	handler := endpointHandler{
		proxy:  proxy,
		router: router,
		log:    log,
	}

	handler.handlerInit()

	return handler.router
}

func (e *endpointHandler) handlerInit() {
	e.router.HandleFunc("/v1/status", e.statusHandler)
	e.router.HandleFunc("/v1/file/{id}", e.fileFetcher).Methods(http.MethodGet)
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

func (e *endpointHandler) fileFetcher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileVar := vars["id"]

	fileID, err := strconv.Atoi(fileVar)
	if err != nil {
		e.log.Println(err)
	}

	file, err := e.proxy.GetFile(fileID)
	if err != nil {
		e.log.Println(err)
	}

	_, err = w.Write(file)
	if err != nil {
		e.log.Println(err)
	}
}
