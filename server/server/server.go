package server

import (
	"github.com/labstack/echo/v4"
	"github.com/soundreaper/portal/db"
	"gorm.io/gorm"
)

// Server is a wrapper around echo server and database connection
type Server struct {
	e  *echo.Echo
	db *gorm.DB
}

// NewServer creates a Server
func NewServer(database *gorm.DB) *Server {
	if database == nil {
		database = db.Connect()
	}

	return &Server{
		e:  echo.New(),
		db: database,
	}
}

// GetDB will return the instance of the DB
func (s *Server) GetDB() *gorm.DB {
	return s.db
}

// Start Server functionality
func (s *Server) Start(port string) {
	// register routes
	s.Routes()
	// start the server
	s.e.Logger.Fatal(s.e.Start(port))
}

// Close stops the Server
func (s *Server) Close() {
	// stop the server
	s.e.Close()
}
