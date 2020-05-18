package application

import (
	"golang-ddd/domain/entity"
	"golang-ddd/domain/repository"
)

// GroupApplication struct that holds the logic layer
type GroupApplication struct {
	repo repository.GroupRepository
}

// GroupAppInterface interface
type GroupAppInterface interface {
	repository.GroupRepository
}

var _ GroupAppInterface = &GroupApplication{}

// All group
func (app GroupApplication) All(perPage int, pageNumber int) (*entity.Groups, error) {
	return app.repo.All(perPage, pageNumber)
}

// Create group
func (app GroupApplication) Create(group entity.Group) (*entity.Group, error) {
	return app.repo.Create(group)
}

// Update group
func (app GroupApplication) Update(id uint, data map[string]interface{}) (*entity.Group, error) {
	return app.repo.Update(id, data)
}

// Delete group
func (app GroupApplication) Delete(id uint) error {
	return app.repo.Delete(id)
}

// Find group
func (app GroupApplication) Find(id uint) (*entity.Group, error) {
	return app.repo.Find(id)
}
