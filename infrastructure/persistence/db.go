package persistence

import (
	"ddd/domain/entity"
	"ddd/domain/repository"
	"ddd/infrastructure/database"

	"github.com/jinzhu/gorm"
)

// Repositories its just like a container holding all register repositories in the app
type Repositories struct {
	User repository.UserRepository
	db   *gorm.DB
}

// NewRepositories : its just like a container holding all register repositories in the app
func NewRepositories() (*Repositories, error) {
	db := database.Connect()

	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}

// Close the  database connection
func (repositories *Repositories) Close() error {
	return repositories.db.Close()
}

// AutoMigrate all tables
func (repositories *Repositories) AutoMigrate() error {
	return repositories.db.AutoMigrate(
		&entity.User{},
	).Error
}
