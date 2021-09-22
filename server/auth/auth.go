package auth

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var (
	err             = errors.New("could not get user from context")
	unAuthorizedMsg = echo.NewHTTPError(http.StatusUnauthorized, "unauthorized to access that resource")
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

// CheckPermissions checks if a given user has persmission to access a resource via the user ID from the JWT token
// It will return true if the user has permission to access the resource, otherwise false
// If the function returns false, it will also return an HTTP error that can be returned from the handler
func CheckPermissions(c echo.Context, userID string) (bool, error) {
	// get UserID using helper function
	authorizedID, err := GetUserIDFromContext(c)

	// failed to get the UserID so assume there is an error and do not let user proceed
	if err != nil {
		return false, unAuthorizedMsg
	}

	// if the user ID does not match the user ID from the JWT token, they cannot access the resource
	if userID != authorizedID {
		return false, unAuthorizedMsg
	}

	// the user passed the check, they can access the resource
	return true, nil
}
