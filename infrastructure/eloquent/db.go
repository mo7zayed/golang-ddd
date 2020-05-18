package eloquent

import (
	"golang-ddd/domain/entity"
	"golang-ddd/domain/repository"
	"golang-ddd/infrastructure/database"

	"github.com/jinzhu/gorm"
)

// Repositories its just like a container holding all register repositories in the app
type Repositories struct {
	User    repository.UserRepository
	Group   repository.GroupRepository
	Message repository.MessageRepository
	db      *gorm.DB
}

// NewRepositories : its just like a container holding all register repositories in the app
func NewRepositories() (*Repositories, error) {
	db := database.Connect()

	return &Repositories{
		User:    NewUserRepository(db),
		Group:   NewGroupRepository(db),
		Message: NewMessageRepository(db),
		db:      db,
	}, nil
}

// Statement method to execute a custom database statement
func (repositories *Repositories) Statement(statement string, values ...interface{}) {
	repositories.db.Exec(statement, values)
}

// Close the  database connection
func (repositories *Repositories) Close() error {
	return repositories.db.Close()
}

// AutoMigrate all tables
func (repositories *Repositories) AutoMigrate() error {
	return repositories.db.AutoMigrate(
		&entity.User{},
		&entity.Group{},
		&entity.Message{},
	).Error
}
