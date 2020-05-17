package entity

import (
	"ddd/infrastructure/models"
)

// Group Entity
type Group struct {
	models.BaseModel
	Name string `gorm:"size:100;not null;" json:"name"`
}

// Groups ...
type Groups []Group
