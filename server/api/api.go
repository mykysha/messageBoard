package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nndergunov/messageBoard/server/pkg/logger"
)

// API is main server handler.
type API struct {
	router *mux.Router
	log    *logger.Logger
}

// NewAPI returns new instance of api.API.
func NewAPI(router *mux.Router, log *logger.Logger) *API {
	api := &API{
		router: router,
		log:    log,
	}

	return api
}

// ServeHTTP method satisfies http.Handler interface.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
