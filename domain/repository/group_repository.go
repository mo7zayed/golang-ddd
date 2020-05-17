package repository

import "ddd/domain/entity"

// GroupRepository instance
type GroupRepository interface {
	All(perPage int, pageNumber int) (*entity.Groups, error)
	Create(user entity.Group) (*entity.Group, error)
	Update(id uint, data map[string]interface{}) (*entity.Group, error)
	Delete(id uint) error
	Find(id uint) (*entity.Group, error)
}
