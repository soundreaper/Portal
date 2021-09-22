package handlers

import (
	"html"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/soundreaper/portal/auth"
	"github.com/soundreaper/portal/models"
)

func (h *Handler) Delete(c echo.Context) error {
	// ID of the image to get
	imageID := html.EscapeString(c.Param("objectID"))
	// Create a new image store to interact with the database
	store := models.NewImageStore(h.db)
	// Get the image from the database
	i, err := store.GetByID(imageID)
	if err != nil {
		log.Println("could not get image with given id. err:", err)
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	// Check permissions before returning the image
	allowed, err := auth.CheckPermissions(c, i.UserID)
	if !allowed {
		return err
	}
	// Delete image with the given ID
	err = store.DeleteByID(imageID)
	if err != nil {
		log.Println("failed to delete image. err:", err)
		return echo.NewHTTPError(http.StatusNotFound, "failed to delete image")
	}
	// Return number of images deleted
	resp := map[string]bool{"success": true}
	return c.JSON(http.StatusOK, resp)
}
