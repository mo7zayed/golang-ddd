package persistence

import (
	"ddd/domain/entity"
	"ddd/domain/repository"
	"ddd/utils/helpers"
	"errors"

	"github.com/jinzhu/gorm"
)

// UserEloquentRepository that holds the common needed methods in the repository
type UserEloquentRepository struct {
	db *gorm.DB
}

//UserRepo implements the repository.UserRepository interface
var _ repository.UserRepository = &UserEloquentRepository{}

// NewUserRepository : Instanciate new repository instance
func NewUserRepository(db *gorm.DB) *UserEloquentRepository {
	return &UserEloquentRepository{db}
}

// Create new user
func (r *UserEloquentRepository) Create(user entity.User) (*entity.User, error) {
	err := r.db.Debug().Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update user
func (r *UserEloquentRepository) Update(id uint, data map[string]interface{}) (*entity.User, error) {
	var user entity.User

	err := r.db.Debug().Model(&user).Updates(data).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Delete user
func (r *UserEloquentRepository) Delete(id uint) error {
	err := r.db.Debug().Where("id LIKE ?", id).Delete(entity.User{}).Error

	if err != nil {
		return err
	}

	return nil
}

// Find a user
func (r *UserEloquentRepository) Find(id uint) (*entity.User, error) {
	var user entity.User

	err := r.db.Debug().First(&user, id).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindUserByEmailAndPassword : finding user by password or email
func (r *UserEloquentRepository) FindUserByEmailAndPassword(email string, password string) (*entity.User, error) {
	var user entity.User

	// Get first matched record
	err := r.db.Debug().Where("email = ?", email).First(&user).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("wrong credentials")
	}

	if err != nil {
		return nil, err
	}

	// Verify the password
	if !helpers.CompareHash(user.Password, password) {
		return nil, errors.New("wrong credentials")
	}

	return &user, nil
}
