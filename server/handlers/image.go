package handlers

/*

import (
	"html"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/soundreaper/portal/auth"
	"github.com/soundreaper/portal/models"
	"github.com/soundreaper/portal/s3"
)

// GetImage will return an image given its ID
func (h *Handler) GetImage(c echo.Context) error {
	// ID of the image to get
	imageID := html.EscapeString(c.Param("imageID"))
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

	return c.JSON(http.StatusOK, i)
}

// GetUserImages will return all images for a given user
func (h *Handler) GetUserImages(c echo.Context) error {
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

// CreateImage will create a new image given image file and user ID
func (h *Handler) CreateImage(c echo.Context) error {
	// Grab user ID using helper function in auth
	userID, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return err
	}
	// Get the image file from form post
	iFile, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "please provide valid image file")
	}
	// Upload image to S3
	bucket := s3.NewS3Session()
	url, fileName, err := bucket.Upload(iFile)
	if err != nil {
		log.Println("failed to upload to S3 bucket. err:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to upload image")
	}
	// Create image model and bind to given user in database
	// Main logic for image handling is in this function
	rec, err := models.CreateUploadImage(h.db, url, fileName, userID)
	if err != nil {
		log.Println("error creating image model. err:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create image in database")
	}

	return c.JSON(http.StatusOK, rec)
}

// DeleteImage will delete an image given the image ID
func (h *Handler) DeleteImage(c echo.Context) error {
	// ID of the image to get
	imageID := html.EscapeString(c.Param("imageID"))
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
	deleted, err := store.DeleteByID(imageID)
	if err != nil {
		log.Println("failed to delete image. err:", err)
		return echo.NewHTTPError(http.StatusNotFound, "could not find image to delete")
	}
	// Return number of images deleted
	resp := map[string]int64{"deleted:": deleted}
	return c.JSON(http.StatusOK, resp)
}
*/
