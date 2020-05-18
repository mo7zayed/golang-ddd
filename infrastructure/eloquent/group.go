package eloquent

import (
	"golang-ddd/domain/entity"
	"golang-ddd/domain/repository"
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

// All groups
func (r *GroupEloquentRepository) All(perPage int, pageNumber int) (*entity.Groups, error) {
	var groups entity.Groups

	err := r.db.Limit(perPage).Offset((pageNumber - 1) * perPage).Find(&groups).Error

	if err != nil {
		return nil, err
	}

	return &groups, nil
}

// Create new group
func (r *GroupEloquentRepository) Create(group entity.Group) (*entity.Group, error) {
	err := r.db.Create(&group).Error

	if err != nil {
		return nil, err
	}

	return &group, nil
}

// Update group
func (r *GroupEloquentRepository) Update(id uint, data map[string]interface{}) (*entity.Group, error) {
	var group entity.Group

	err := r.db.First(&group, id).Updates(data).Error

	if err != nil {
		return nil, err
	}

	return &group, nil
}

// Delete group
func (r *GroupEloquentRepository) Delete(id uint) error {
	err := r.db.Where("id LIKE ?", id).Delete(entity.Group{}).Error

	if err != nil {
		return err
	}

	return nil
}

// Find a group
func (r *GroupEloquentRepository) Find(id uint) (*entity.Group, error) {
	var group entity.Group

	err := r.db.First(&group, id).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("group not found")
	}

	if err != nil {
		return nil, err
	}

	return &group, nil
}
