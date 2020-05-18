package entity

import (
	"golang-ddd/infrastructure/models"
)

// Message Entity
type Message struct {
	models.BaseModel
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
	Channel  string `json:"channel"`
	Message  string `json:"message"`
}


// Messages ...
type Messages []Message
