package helpers

import "github.com/labstack/echo"

// Respond Function Is To Handel Http Returns With A Json.
func Respond(c echo.Context, status int, success bool, payload map[string]interface{}) error {
	return c.JSON(status, map[string]interface{}{
		"status":  status,
		"success": success,
		"payload": payload,
	})
}
