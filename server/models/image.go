package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Image represents an image in the database
type Image struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	UserID     string    `json:"-"`
	URL        string    `json:"url"`
	UploadName string    `json:"upload_name"`
	UploadDate time.Time `json:"upload_date"`
}

// ImageStore represents a way to interact with image models
type ImageStore struct {
	db *gorm.DB
}

// NewImageStore will create a new image store with the given database
func NewImageStore(db *gorm.DB) *ImageStore {
	return &ImageStore{
		db,
	}
}

// Create will create a new image in the database (this function is for manual image creation only)
func (i *ImageStore) Create(image Image) (*Image, error) {
	err := i.db.Create(&image).Error
	if err != nil {
		return &Image{}, err
	}

	return &image, nil
}

func CreateUploadImage(db *gorm.DB, url string, filename string, userID string) (*Image, error) {
	userStore := NewUserStore(db)

	user, err := userStore.GetByID(userID)
	if err != nil {
		return &Image{}, err
	}

	image := Image{
		URL:        url,
		UserID:     user.ID,
		UploadName: filename,
		UploadDate: time.Now(),
	}

	err = db.Create(&image).Error
	if err != nil {
		return &Image{}, err
	}

	return &image, nil
}

// GetByID will get an image by its ID
func (i *ImageStore) GetByID(id string) (*Image, error) {
	var image Image
	err := i.db.Model(&Image{}).Where("id = ?", id).Take(&image).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Image{}, errors.New("image: image not found")
		}
		return &Image{}, err
	}
	return &image, nil
}

func (i *ImageStore) GetImagesByUser(id string) ([]Image, error) {
	var images []Image
	err := i.db.Where("name <> ?", id).Find(&images).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []Image{}, errors.New("user: user not found")
		}

		return []Image{}, err
	}
	return images, nil
}

// DeleteByID will delete an image by its ID
func (i *ImageStore) DeleteByID(imageID string) error {
	db := i.db.Delete(&Image{}, Image{ID: imageID})
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return errors.New("image not found")
		}
		return db.Error
	}
	return nil
}

// BeforeCreate is called before an image in created in the database (GORM Hook)
func (i *Image) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ID == "" {
		i.ID = uuid.New().String()
	}

	return
}
