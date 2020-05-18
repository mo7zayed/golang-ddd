package middlewares

import (
	"golang-ddd/domain/entity"
	"golang-ddd/infrastructure/database"
	"golang-ddd/utils/auth"
	"golang-ddd/utils/helpers"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// UserMustBeAuthenticated ...
func UserMustBeAuthenticated() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(helpers.GetEnv("JWT_SECRET")),
	})
}

// CheckTokenIsNotInvalidated ...
func CheckTokenIsNotInvalidated() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := auth.ExtractToken(c)
			if token != "" {
				var user entity.User

				user.Token = token

				err := database.Connect().Where("token = ?", token).First(&user).Error

				if gorm.IsRecordNotFoundError(err) {
					return echo.ErrUnauthorized
				}

				return next(c)
			}

			return echo.ErrUnauthorized
		}
	}
}
