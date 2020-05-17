package interfaces

import (
	"ddd/application"
	"ddd/domain/entity"
	"ddd/utils/auth"
	"ddd/utils/helpers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

// AuthHandler ...
type AuthHandler struct {
	user application.UserAppInterface
}

// NewAuthHandler Handler Instance
func NewAuthHandler(user application.UserAppInterface) *AuthHandler {
	return &AuthHandler{
		user: user,
	}
}

// Login method
func (handler *AuthHandler) Login(c echo.Context) error {
	errors := govalidator.New(govalidator.Options{
		Request: c.Request(), // request object
		Rules: govalidator.MapData{
			"email":    []string{"required", "max:90", "email"},
			"password": []string{"required", "min:6", "max:100"},
		},
		RequiredDefault: true, // all the field to be pass the rules
	}).Validate()

	if len(errors) > 0 {
		return helpers.RespondValidationErrors(c, errors)
	}

	user, err := handler.user.FindUserByEmailAndPassword(c.FormValue("email"), c.FormValue("password"))

	if err != nil {
		return helpers.RespondValidationErrors(c, map[string][]string{
			"email": {
				"These cerdentials dosen't match our records",
			},
		})
	}

	data, err := user.GenerateToken()

	token, ok := data["token"]

	if err != nil || !ok {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Error generating user token",
		}, http.StatusInternalServerError, false)
	}

	// update the token field in database.
	handler.user.Update(user.ID, map[string]interface{}{
		"token": token,
	})

	return helpers.Respond(c, data, http.StatusOK, true)
}

// Register method registers a new user
func (handler *AuthHandler) Register(c echo.Context) error {
	errors := govalidator.New(govalidator.Options{
		Request: c.Request(), // request object
		Rules: govalidator.MapData{
			"name":     []string{"required", "min:2", "max:50"},
			"email":    []string{"required", "max:90", "email"},
			"password": []string{"required", "min:6", "max:100"},
		},
		RequiredDefault: true, // all the field to be pass the rules
	}).Validate()

	if len(errors) > 0 {
		return helpers.RespondValidationErrors(c, errors)
	}

	user, _ := handler.user.FindUserByEmail(c.FormValue("email"))

	if user != nil {
		return helpers.RespondValidationErrors(c, map[string][]string{
			"email": {
				"Email already taken",
			},
		})
	}

	preparedUser := entity.User{}
	preparedUser.Name = c.FormValue("name")
	preparedUser.Email = c.FormValue("email")
	preparedUser.Password = c.FormValue("password")

	user, err := handler.user.Create(preparedUser)

	if err != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Error while creating new user",
		}, http.StatusInternalServerError, false)
	}

	data, err := user.GenerateToken()

	if err != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Error generating user token",
		}, http.StatusInternalServerError, false)
	}

	return helpers.Respond(c, data, http.StatusOK, true)
}

// Me get the current authinticated user data
func (handler *AuthHandler) Me(c echo.Context) error {
	user, err := auth.User(c)

	if err != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": err.Error,
		}, http.StatusInternalServerError, false)
	}

	return helpers.Respond(c, map[string]interface{}{
		"user": user,
	}, http.StatusOK, true)
}

// Logout logs the user out and invalidate his token
func (handler *AuthHandler) Logout(c echo.Context) error {
	token := auth.ExtractToken(c)

	if token != "" {
		user, err := handler.user.FindUserByToken(token)

		if err != nil {
			return helpers.Respond(c, map[string]interface{}{
				"message": "Error Finding user by token",
			}, http.StatusInternalServerError, false)
		}

		// update the token field in database.
		handler.user.Update(user.ID, map[string]interface{}{
			"token": nil,
		})
	}

	return helpers.Respond(c, map[string]interface{}{
		"message": "logged out successfully",
	}, http.StatusOK, true)
}

// Refresh user token
func (handler *AuthHandler) Refresh(c echo.Context) error {
	token := auth.ExtractToken(c)

	user, _ := handler.user.FindUserByToken(token)

	data, err := user.GenerateToken()

	newToken, ok := data["token"]

	if err != nil || !ok {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Error generating user token",
		}, http.StatusInternalServerError, false)
	}

	// update the token field in database.
	handler.user.Update(user.ID, map[string]interface{}{
		"token": newToken,
	})

	return helpers.Respond(c, data, http.StatusOK, true)
}
