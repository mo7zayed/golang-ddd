package entity

import (
	"golang-ddd/infrastructure/models"
	"golang-ddd/utils/helpers"
	"html"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// User Entity
type User struct {
	models.BaseModel
	Name     string `gorm:"size:100;not null;" json:"name"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
	Token    string `gorm:"size:500;null;" json:"token"`
	UserType string `gorm:"size:10;null;" sql:"DEFAULT:'user'" json:"user_type"`
}

// PublicUser instance that will be displayable
type PublicUser struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
}

// Users ...
type Users []User

// BeforeSave is a gorm hook
func (u *User) BeforeSave() error {
	u.Password = helpers.Hash(u.Password)
	return nil
}

// PublicUser : So that we dont expose the user's email address and password to the world
func (u *User) PublicUser() PublicUser {
	return PublicUser{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		UserType:  u.UserType,
		CreatedAt: u.CreatedAt,
	}
}

// Prepare the user to be saved in the database
func (u *User) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// GenerateToken From User
func (u User) GenerateToken() (map[string]interface{}, error) {
	// jwt claims
	// To add other keys check jwt.StandardClaims first
	claims := jwt.MapClaims{}
	claims["user"] = u.PublicUser()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString(
		[]byte(helpers.GetEnv("JWT_SECRET")),
	)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token":  accessToken,
		"claims": claims,
	}, nil
}

// PublicUsers : So that we dont expose the user's email address and password to the world
func (users Users) PublicUsers() []PublicUser {
	result := make([]PublicUser, len(users))
	for index, user := range users {
		result[index] = user.PublicUser()
	}
	return result
}
