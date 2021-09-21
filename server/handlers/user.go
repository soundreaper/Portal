package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/soundreaper/portal/auth"
	"github.com/soundreaper/portal/models"
)

// GetUser is the handler for getting a user
func (h *Handler) GetUser(c echo.Context) error {
	// Grab user ID using helper function in auth
	userID, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// setup user store
	store := models.NewUserStore(h.db)
	// get the user with given ID
	user, err := store.GetByID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

// CreateUser is the handler for creating a new user with the passed body
func (h *Handler) CreateUser(c echo.Context) error {
	// Grab user ID using helper function in auth
	userID, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// setup user store
	store := models.NewUserStore(h.db)
	// create new user object to bind body to
	u := models.User{}
	// attempt to bind body to user object
	err = c.Bind(&u)
	if err != nil {
		// return error if binding fails
		return echo.NewHTTPError(http.StatusBadRequest, "bad body data")
	}
	// check to see if the user already exists, no dupes!
	user, _ := store.GetByID(userID)
	if user.ID != "" {
		log.Printf("user details for the id %s already exists\n", userID)
		return echo.NewHTTPError(http.StatusBadRequest, "unauthorized")
	}

	// set user ID to ID attached to auth from context
	u.ID = userID

	// save user to the database
	user, err = store.Create(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to create user with provided details")
	}

	// return basic user data
	return c.JSON(http.StatusCreated, user)
}

// DeleteUser is the handler for deleting a user given the user ID
func (h *Handler) DeleteUser(c echo.Context) error {
	// Grab user ID using helper function in auth
	userID, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// setup user store
	store := models.NewUserStore(h.db)
	// delete the user with given ID
	_, err = store.DeleteByID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to delete user")
	}

	r := map[string]string{"msg": "deleted user"}
	return c.JSON(http.StatusOK, r)
}
