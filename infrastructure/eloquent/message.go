package eloquent

import (
	"ddd/domain/entity"
	"ddd/domain/repository"

	"github.com/jinzhu/gorm"
)

// MessageEloquentRepository that holds the common needed methods in the repository
type MessageEloquentRepository struct {
	db *gorm.DB
}

// MessageRepo implements the repository.MessageRepository interface
var _ repository.MessageRepository = &MessageEloquentRepository{}

// NewMessageRepository : Instanciate new repository instance
func NewMessageRepository(db *gorm.DB) *MessageEloquentRepository {
	return &MessageEloquentRepository{db}
}

// ListByChannelName ...
func (r *MessageEloquentRepository) ListByChannelName(channelName string) (*entity.Messages, error) {
	var messages entity.Messages

	err := r.db.Where("channel=?", channelName).Find(&messages).Error

	if err != nil {
		return nil, err
	}

	return &messages, nil
}

// Create new message
func (r *MessageEloquentRepository) Create(message entity.Message) (*entity.Message, error) {
	err := r.db.Create(&message).Error

	if err != nil {
		return nil, err
	}

	return &message, nil
}
