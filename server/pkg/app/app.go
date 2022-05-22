package app

import (
	"fmt"

	"github.com/nndergunov/messageBoard/server/pkg/domain"
)

type App struct {
	storage Storage
}

func NewApp(storage Storage) *App {
	return &App{storage: storage}
}

func (a App) ReturnAllMessages() ([]domain.Message, error) {
	messages, err := a.storage.ReadAllMessages()
	if err != nil {
		return nil, fmt.Errorf("getting messages from storage: %w", err)
	}

	return messages, nil
}

func (a App) InsertMessage(message domain.Message) error {
	_, err := a.storage.InsertUser(message.Author)
	if err != nil {
		return fmt.Errorf("inserting author to the storage: %w", err)
	}

	err = a.storage.InsertMessage(message)
	if err != nil {
		return fmt.Errorf("inserting message to the storage: %w", err)
	}

	return nil
}
