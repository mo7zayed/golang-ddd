package entity

import (
	"ddd/infrastructure/models"
	"ddd/utils/helpers"
	"html"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/thedevsaddam/govalidator"
)

// User Entity
type User struct {
	models.BaseModel
	Name     string `gorm:"size:100;not null;" json:"first_name"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
}

// PublicUser instance that will be displayable
type PublicUser struct {
	ID        uint      `json:"id"`
	Name      string    `gorm:"size:100;not null;" json:"first_name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
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

// Validate Entity Data
func (u *User) Validate(request *http.Request) map[string][]string {

	opts := govalidator.Options{
		Request: request, // request object
		Rules: govalidator.MapData{
			"name":     []string{"required", "max:90"},
			"email":    []string{"required", "max:90", "email"},
			"password": []string{"required", "min:6", "max:100"},
		},
		RequiredDefault: true, // all the field to be pass the rules
	}

	return govalidator.New(opts).Validate()
}

// GenerateToken From User
func (u User) GenerateToken() (map[string]interface{}, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = u
	claims["expiration"] = time.Now().Add(time.Hour * 72).Unix()

	accessToken, err := token.SignedString([]byte(helpers.Hash(helpers.GetEnv("JWT_SECRET"))))

	if err != nil {
		return map[string]interface{}{}, err
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
