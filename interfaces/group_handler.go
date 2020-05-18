package interfaces

import (
	"golang-ddd/application"
	"golang-ddd/domain/entity"
	"golang-ddd/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

// GroupsHandler struct defines the dependencies that will be used
type GroupsHandler struct {
	group application.GroupAppInterface
}

// NewGroupHandler instance from group handler.
func NewGroupHandler(group application.GroupAppInterface) *GroupsHandler {
	return &GroupsHandler{
		group: group,
	}
}

// Index method
func (handler *GroupsHandler) Index(c echo.Context) error {
	perPage := 20

	pageNumber, _ := strconv.Atoi(c.FormValue("page"))

	groups, err := handler.group.All(perPage, pageNumber)

	if err != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": err.Error,
		}, http.StatusInternalServerError, false)
	}

	return helpers.Respond(c, map[string]interface{}{
		"data":  groups,
		"limit": perPage,
		"page":  pageNumber,
	}, http.StatusOK, true)
}

// Create method
func (handler *GroupsHandler) Create(c echo.Context) error {
	errors := govalidator.New(govalidator.Options{
		Request: c.Request(), // request object
		Rules: govalidator.MapData{
			"name": []string{"required", "min:2", "max:50"},
		},
		RequiredDefault: true, // all the field to be pass the rules
	}).Validate()

	if len(errors) > 0 {
		return helpers.RespondValidationErrors(c, errors)
	}

	preparedGroup := entity.Group{}
	preparedGroup.Name = c.FormValue("name")

	group, err := handler.group.Create(preparedGroup)

	if err != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Error while creating new group",
		}, http.StatusInternalServerError, false)
	}

	return helpers.Respond(c, map[string]interface{}{
		"message": "Group Created Successfully",
		"data":    group,
	}, http.StatusOK, true)
}

// Show method
func (handler *GroupsHandler) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	group, err := handler.group.Find(uint(id))

	if err != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Entity Not Found",
		}, http.StatusNotFound, false)
	}

	return helpers.Respond(c, map[string]interface{}{
		"data": group,
	}, http.StatusOK, true)
}

// Update method
func (handler *GroupsHandler) Update(c echo.Context) error {
	errors := govalidator.New(govalidator.Options{
		Request: c.Request(), // request object
		Rules: govalidator.MapData{
			"name": []string{"required", "min:2", "max:50"},
		},
		RequiredDefault: true, // all the field to be pass the rules
	}).Validate()

	if len(errors) > 0 {
		return helpers.RespondValidationErrors(c, errors)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	group, err := handler.group.Update(uint(id), map[string]interface{}{
		"name": c.FormValue("name"),
	})

	if err != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Error while updating group",
		}, http.StatusInternalServerError, false)
	}

	return helpers.Respond(c, map[string]interface{}{
		"message": "Group Updated Successfully",
		"data":    group,
	}, http.StatusOK, true)
}

// Delete method
func (handler *GroupsHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	group, err := handler.group.Find(uint(id))

	if err != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Entity Not Found",
		}, http.StatusNotFound, false)
	}

	if handler.group.Delete(group.ID) != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Error deleting group",
		}, http.StatusNotFound, false)
	}

	return helpers.Respond(c, map[string]interface{}{
		"message": "Group deleted successfully",
	}, http.StatusOK, true)
}
