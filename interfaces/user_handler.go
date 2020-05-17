package interfaces

import (
	"ddd/application"
	"ddd/domain/entity"
	"ddd/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

// UsersHandler struct defines the dependencies that will be used
type UsersHandler struct {
	user application.UserAppInterface
}

// NewUserHandler instance from user handler.
func NewUserHandler(user application.UserAppInterface) *UsersHandler {
	return &UsersHandler{
		user: user,
	}
}

// Index method
func (handler *UsersHandler) Index(c echo.Context) error {
	perPage := 2000

	pageNumber, _ := strconv.Atoi(c.FormValue("page"))

	users, err := handler.user.All(perPage, pageNumber)

	if err != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": err.Error,
		}, http.StatusInternalServerError, false)
	}

	return helpers.Respond(c, map[string]interface{}{
		"data":  users.PublicUsers(),
		"limit": perPage,
		"page":  pageNumber,
	}, http.StatusOK, true)
}

// Create method
func (handler *UsersHandler) Create(c echo.Context) error {
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

	return helpers.Respond(c, map[string]interface{}{
		"message": "User Created Successfully",
		"data":    user.PublicUser(),
	}, http.StatusOK, true)
}

// Show method
func (handler *UsersHandler) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := handler.user.Find(uint(id))

	if err != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Entity Not Found",
		}, http.StatusNotFound, false)
	}

	return helpers.Respond(c, map[string]interface{}{
		"data": user.PublicUser(),
	}, http.StatusOK, true)
}

// Update method
func (handler *UsersHandler) Update(c echo.Context) error {
	errors := govalidator.New(govalidator.Options{
		Request: c.Request(), // request object
		Rules: govalidator.MapData{
			"name":  []string{"required", "min:2", "max:50"},
			"email": []string{"required", "max:90", "email"},
		},
		RequiredDefault: true, // all the field to be pass the rules
	}).Validate()

	if len(errors) > 0 {
		return helpers.RespondValidationErrors(c, errors)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	emailCheck, _ := handler.user.FindUserByEmail(c.FormValue("email"))

	if emailCheck != nil && (*emailCheck).ID != uint(id) {
		return helpers.RespondValidationErrors(c, map[string][]string{
			"email": {
				"Email already taken",
			},
		})
	}

	user, err := handler.user.Update(uint(id), map[string]interface{}{
		"name":  c.FormValue("name"),
		"email": c.FormValue("email"),
	})

	if err != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Error while updating user",
		}, http.StatusInternalServerError, false)
	}

	return helpers.Respond(c, map[string]interface{}{
		"message": "User Updated Successfully",
		"data":    user.PublicUser(),
	}, http.StatusOK, true)
}

// Delete method
func (handler *UsersHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := handler.user.Find(uint(id))

	if err != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Entity Not Found",
		}, http.StatusNotFound, false)
	}

	if handler.user.Delete(user.ID) != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Error deleting user",
		}, http.StatusNotFound, false)
	}

	return helpers.Respond(c, map[string]interface{}{
		"message": "User deleted successfully",
	}, http.StatusOK, true)
}
