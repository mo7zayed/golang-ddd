package repository

import "ddd/domain/entity"

// UserRepository instance
type UserRepository interface {
	Create(user entity.User) (*entity.User, error)
	Update(id uint, data map[string]interface{}) (*entity.User, error)
	Delete(id uint) error
	Find(id uint) (*entity.User, error)
	FindUserByEmailAndPassword(email string, password string) (*entity.User, error)
}
