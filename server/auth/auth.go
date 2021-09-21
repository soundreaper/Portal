package auth

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var (
	err = errors.New("could not get user from context")
)

// GetUserIDFromContext will return the UserID from context
func GetUserIDFromContext(c echo.Context) (string, error) {
	token := c.Request().Context().Value("user")

	// if we can't get the UserID
	if token == nil {
		return "", err
	}

	// convert the token into a JWT token and then turn the claims into a Golang map
	t := token.(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)

	// get the user ID from claims map
	id, ok := claims["sub"]
	if !ok {
		return "", errors.New("could not get user id from context")
	}

	// return the user ID and no error if we could grab it from claims
	return id.(string), nil
}
