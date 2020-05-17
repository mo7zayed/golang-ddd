package eloquent

import (
	"ddd/domain/entity"
	"ddd/domain/repository"
	"errors"

	"github.com/jinzhu/gorm"
)

// GroupEloquentRepository that holds the common needed methods in the repository
type GroupEloquentRepository struct {
	db *gorm.DB
}

// GroupRepo implements the repository.GroupRepository interface
var _ repository.GroupRepository = &GroupEloquentRepository{}

// NewGroupRepository : Instanciate new repository instance
func NewGroupRepository(db *gorm.DB) *GroupEloquentRepository {
	return &GroupEloquentRepository{db}
}

// All users
func (r *GroupEloquentRepository) All(perPage int, pageNumber int) (*entity.Groups, error) {
	var users entity.Groups

	err := r.db.Limit(perPage).Offset((pageNumber - 1) * perPage).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}

// Create new user
func (r *GroupEloquentRepository) Create(user entity.Group) (*entity.Group, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update user
func (r *GroupEloquentRepository) Update(id uint, data map[string]interface{}) (*entity.Group, error) {
	var user entity.Group

	err := r.db.First(&user, id).Updates(data).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Delete user
func (r *GroupEloquentRepository) Delete(id uint) error {
	err := r.db.Where("id LIKE ?", id).Delete(entity.Group{}).Error

	if err != nil {
		return err
	}

	return nil
}

// Find a user
func (r *GroupEloquentRepository) Find(id uint) (*entity.Group, error) {
	var user entity.Group

	err := r.db.First(&user, id).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
