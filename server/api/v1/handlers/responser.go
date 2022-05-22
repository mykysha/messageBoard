package handlers

import (
	"net/http"

	v1 "github.com/nndergunov/messageBoard/server/api/v1"
)

func (e endpointHandler) respond(data any, responseWriter http.ResponseWriter) {
	encodedData, err := v1.Encode(data)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = responseWriter.Write(encodedData)
	if err != nil {
		e.log.Println(err)

		return
	}
}
