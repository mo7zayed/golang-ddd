package interfaces

import (
	"ddd/application"
	"ddd/utils/helpers"
	"net/http"

	"github.com/labstack/echo"
)

// Users struct defines the dependencies that will be used
type UsersHandler struct {
	app application.UserAppInterface
}

// NewUsersHandler instance from user handler.
func NewUserHandler(app application.UserAppInterface) *UsersHandler {
	return &UsersHandler{
		app: app,
	}
}

// Index method
func (handler *UsersHandler) Index(c echo.Context) error {
	return helpers.Respond(c, map[string]interface{}{
		"data": "here",
	}, http.StatusOK, true)
}

// Create method
func (handler *UsersHandler) Create(c echo.Context) error {
	return helpers.Respond(c, map[string]interface{}{
		"data": "here",
	}, http.StatusOK, true)
}

// Show method
func (handler *UsersHandler) Show(c echo.Context) error {
	return helpers.Respond(c, map[string]interface{}{
		"data": "here",
	}, http.StatusOK, true)
}

// Update method
func (handler *UsersHandler) Update(c echo.Context) error {
	return helpers.Respond(c, map[string]interface{}{
		"data": "here",
	}, http.StatusOK, true)
}

// Delete method
func (handler *UsersHandler) Delete(c echo.Context) error {
	return helpers.Respond(c, map[string]interface{}{
		"data": "here",
	}, http.StatusOK, true)
}
