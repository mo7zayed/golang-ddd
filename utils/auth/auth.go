package auth

import (
	"fmt"
	"golang-ddd/utils/helpers"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// TokenHolderName : which request header holds the token ?
var TokenHolderName string = "Authorization"

// TokenValid : checks if the token is valid
func TokenValid(c echo.Context) error {
	token, err := VerifyToken(c)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

// VerifyToken : checks if the token format is created by our secret
func VerifyToken(c echo.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(helpers.GetEnv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractToken from the request
func ExtractToken(c echo.Context) string {
	keys := c.Request().URL.Query()

	token := keys.Get("token")

	if token != "" {
		return token
	}

	bearToken := c.Request().Header.Get(TokenHolderName)

	if arr := strings.Split(bearToken, " "); len(arr) == 2 {
		return arr[1]
	}

	return ""
}

// User gets the current authineticated user
func User(c echo.Context) (interface{}, error) {
	token, err := VerifyToken(c)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims) // the token claims should conform to MapClaims

	if ok && token.Valid {
		user, ok := claims["user"]

		if !ok {
			return nil, err
		}

		return user, nil
	}

	return nil, err
}
