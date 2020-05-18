package application

import (
	"ddd/domain/entity"
	"ddd/domain/repository"
)

// MessageApplication struct that holds the logic layer
type MessageApplication struct {
	repo repository.MessageRepository
}

// MessageAppInterface interface
type MessageAppInterface interface {
	repository.MessageRepository
}

var _ MessageAppInterface = &MessageApplication{}

// ListByChannelName lists messages by channel name
func (app MessageApplication) ListByChannelName(channelName string) (*entity.Messages, error) {
	return app.repo.ListByChannelName(channelName)
}

// Create message
func (app MessageApplication) Create(message entity.Message) (*entity.Message, error) {
	return app.repo.Create(message)
}
