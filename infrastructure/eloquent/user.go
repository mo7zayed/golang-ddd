package eloquent

import (
	"golang-ddd/domain/entity"
	"golang-ddd/domain/repository"
	"golang-ddd/utils/helpers"
	"errors"

	"github.com/jinzhu/gorm"
)

// UserEloquentRepository that holds the common needed methods in the repository
type UserEloquentRepository struct {
	db *gorm.DB
}

// UserRepo implements the repository.UserRepository interface
var _ repository.UserRepository = &UserEloquentRepository{}

// NewUserRepository : Instanciate new repository instance
func NewUserRepository(db *gorm.DB) *UserEloquentRepository {
	return &UserEloquentRepository{db}
}

// All users
func (r *UserEloquentRepository) All(perPage int, pageNumber int) (*entity.Users, error) {
	var users entity.Users

	err := r.db.Limit(perPage).Offset((pageNumber - 1) * perPage).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}

// Create new user
func (r *UserEloquentRepository) Create(user entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update user
func (r *UserEloquentRepository) Update(id uint, data map[string]interface{}) (*entity.User, error) {
	var user entity.User

	err := r.db.First(&user, id).Updates(data).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Delete user
func (r *UserEloquentRepository) Delete(id uint) error {
	err := r.db.Where("id LIKE ?", id).Delete(entity.User{}).Error

	if err != nil {
		return err
	}

	return nil
}

// Find a user
func (r *UserEloquentRepository) Find(id uint) (*entity.User, error) {
	var user entity.User

	err := r.db.First(&user, id).Error

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
	err := r.db.Where("email = ?", email).First(&user).Error

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

// FindUserByEmail : finding user by email
func (r *UserEloquentRepository) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	// Get first matched record
	err := r.db.Where("email = ?", email).First(&user).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("wrong credentials")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindUserByToken : finding user by token
func (r *UserEloquentRepository) FindUserByToken(token string) (*entity.User, error) {
	var user entity.User

	// Get first matched record
	err := r.db.Where("token = ?", token).First(&user).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("wrong credentials")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
