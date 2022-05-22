package app

import "github.com/nndergunov/messageBoard/server/pkg/domain"

type Storage interface {
	InsertUser(string) (int, error)
	InsertMessage(domain.Message) error
	ReadAllMessages() ([]domain.Message, error)
}
