package repository

import "ddd/domain/entity"

// UserRepository instance
type UserRepository interface {
	All(perPage int, pageNumber int) (*entity.Users, error)
	Create(user entity.User) (*entity.User, error)
	Update(id uint, data map[string]interface{}) (*entity.User, error)
	Delete(id uint) error
	Find(id uint) (*entity.User, error)
	FindUserByEmailAndPassword(email string, password string) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	FindUserByToken(token string) (*entity.User, error)
}
