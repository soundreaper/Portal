package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

// User represents a database entry of a user
type User struct {
	ID     string  `gorm:"primaryKey" json:"-"`
	Images []Image `gorm:"foreignKey:UserID" json:"images"`
}

// UserStore is used to interact with the user model in the database
type UserStore struct {
	db *gorm.DB
}

// NewUserStore will create a new user store
func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db,
	}
}

// Create will create a new user in the databased based on passed user model
func (u *UserStore) Create(user User) (*User, error) {
	// see if the user is in the database first
	exists, _ := u.GetByID(user.ID)
	// if the user exists and we have no error
	if exists.ID != "" {
		return &User{}, errors.New("user with that id already exists")
	}

	// create the user
	err := u.db.Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

// AppendImage will append the given image to the user with the given ID
func (u *UserStore) AppendImage(id string, image Image) (*User, *Image, error) {
	// get the user
	user, err := u.GetByID(id)
	if err != nil {
		return &User{}, &Image{}, nil
	}

	err = u.db.Model(&user).Association("Receipts").Append(&image)
	if err != nil {
		log.Println("user: error getting images for user. err: ", err)
	}

	user, err = u.GetByID(id)
	if err != nil {
		return &User{}, &Image{}, nil
	}

	// return the updated image
	return user, &image, nil
}

// GetByID returns a user from the db that relates to the given id
func (u *UserStore) GetByID(id string) (*User, error) {
	var user User
	err := u.db.Model(&User{}).Where("id = ?", id).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &User{}, errors.New("user: user not found")
		}

		return &User{}, err
	}

	// get user image relations
	user.Images = []Image{}
	err = u.db.Model(&user).Association("Receipts").Find(&user.Images)
	if err != nil {
		log.Println("user: error getting images for user. err: ", err)
		return &User{}, nil
	}

	return &user, nil
}

// DeleteByID will delete a user with the given id
func (u *UserStore) DeleteByID(id string) (int64, error) {
	db := u.db.Model(&User{}).Where("id = ?", id).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
