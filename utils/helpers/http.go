package helpers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Respond Function Is To Handel Http Returns With A Json.
func Respond(c echo.Context, payload map[string]interface{}, status int, success bool) error {
	return c.JSON(status, map[string]interface{}{
		"status":  status,
		"success": success,
		"payload": payload,
	})
}

// RespondValidationErrors Function Is To Handel Http Returns With A Json.
func RespondValidationErrors(c echo.Context, errors map[string][]string) error {
	return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
		"status":  http.StatusUnprocessableEntity,
		"success": false,
		"payload": map[string]string{
			"message": "Validation Errors Please Check Your Inputs",
		},
		"errors": errors,
	})
}
