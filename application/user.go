package application

import (
	"ddd/domain/entity"
	"ddd/domain/repository"
)

// UserApplication struct that holds the logic layer
type UserApplication struct {
	repo repository.UserRepository
}

// UserAppInterface interface
type UserAppInterface interface {
	repository.UserRepository
}

var _ UserAppInterface = &UserApplication{}

// All user
func (app UserApplication) All(perPage int, pageNumber int) (*entity.Users, error) {
	return app.repo.All(perPage, pageNumber)
}

// Create user
func (app UserApplication) Create(user entity.User) (*entity.User, error) {
	return app.repo.Create(user)
}

// Update user
func (app UserApplication) Update(id uint, data map[string]interface{}) (*entity.User, error) {
	return app.repo.Update(id, data)
}

// Delete user
func (app UserApplication) Delete(id uint) error {
	return app.repo.Delete(id)
}

// Find user
func (app UserApplication) Find(id uint) (*entity.User, error) {
	return app.repo.Find(id)
}

// FindUserByEmailAndPassword usually used to authinticate the user
func (app UserApplication) FindUserByEmailAndPassword(email string, password string) (*entity.User, error) {
	return app.repo.FindUserByEmailAndPassword(email, password)
}

// FindUserByEmail usually used to authinticate the user
func (app UserApplication) FindUserByEmail(email string) (*entity.User, error) {
	return app.repo.FindUserByEmail(email)
}

// FindUserByToken ...
func (app UserApplication) FindUserByToken(token string) (*entity.User, error) {
	return app.repo.FindUserByToken(token)
}
