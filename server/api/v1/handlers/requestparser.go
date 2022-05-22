package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	v1 "github.com/nndergunov/messageBoard/server/api/v1"
	"github.com/nndergunov/messageBoard/server/pkg/domain"
)

func requestToMessage(req *http.Request) (*domain.Message, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body failed: %w", err)
	}

	messageData := new(v1.Message)

	err = v1.Decode(body, messageData)
	if err != nil {
		return nil, fmt.Errorf("reading body failed: %w", err)
	}

	return &domain.Message{
		Author: messageData.Author,
		Text:   messageData.Text,
		Time:   messageData.Time,
	}, nil
}
