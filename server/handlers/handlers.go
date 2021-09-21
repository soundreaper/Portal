package handlers

import (
	"gorm.io/gorm"
)

// Handler is a struct that all handlers are methods of
type Handler struct {
	db *gorm.DB
}

// NewHandler will return a new handler struct with the given DB
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
	}
}
