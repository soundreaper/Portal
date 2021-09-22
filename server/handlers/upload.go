package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/soundreaper/portal/auth"
	"github.com/soundreaper/portal/models"
	"github.com/soundreaper/portal/s3"
)

// Images handler returns all images on a given user's account
func (h *Handler) Upload(c echo.Context) error {
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
