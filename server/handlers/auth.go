package handlers

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	echogothic "github.com/nabowler/echo-gothic"
)

// Handles sending back all user data
func (h *Handler) GetCallback(c echo.Context) error {
	user, err := echogothic.CompleteUserAuth(c)
	if err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}

	log.Println(user)

	return c.JSON(200, map[string]goth.User{"auth": user})
}

// Handles user login
func (h *Handler) Login(c echo.Context) error {
	// Try to get the user without re-authenticating
	if gothUser, err := echogothic.CompleteUserAuth(c); err == nil {
		return c.JSON(200, map[string]goth.User{"user": gothUser})
	} else {
		echogothic.BeginAuthHandler(c)
	}

	return nil
}

// Handles user logout
func (h *Handler) Logout(c echo.Context) error {
	// Logout the user
	echogothic.Logout(c)

	// Should redirect to frontend
	return c.JSON(200, map[string]string{"msg": "sucess"})
}
