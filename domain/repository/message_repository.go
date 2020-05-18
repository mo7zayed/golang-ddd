package repository

import "golang-ddd/domain/entity"

// MessageRepository instance
type MessageRepository interface {
	ListByChannelName(channelName string) (*entity.Messages, error)
	Create(message entity.Message) (*entity.Message, error)
}
