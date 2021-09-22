package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/soundreaper/portal/auth"
	"github.com/soundreaper/portal/models"
)

// Images returns all images on a given user's account
func (h *Handler) Images(c echo.Context) error {
	// Grab user ID using helper function in auth
	userID, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return err
	}
	// Create a new user store to interact with the database
	store := models.NewImageStore(h.db)
	// Get the images from the database
	images, err := store.GetImagesByUser(userID)
	if err != nil {
		log.Println("failed to get images for user id", err)
		return echo.NewHTTPError(http.StatusNotFound, "no images found")
	}

	return c.JSON(http.StatusOK, images)
}
