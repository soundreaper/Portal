package server

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/soundreaper/portal/handlers"
)

func (s *Server) Routes() {
	// Setup Middleware
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))

	// setup our route handlers
	h := handlers.NewHandler(s.db)

	// hello route for a status check
	s.e.GET("/status", h.Hello)

	// Group Under API V1 in case we want to change in the future
	v1 := s.e.Group("/api/v1")

	v1.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("test")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("secret123")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	// User routes
	v1.GET("/user", h.GetUser)
	v1.POST("/user", h.CreateUser)
	v1.DELETE("/user", h.DeleteUser)

	// Image routes
	v1.GET("/user/images", h.GetUserImages)

	v1.GET("/image/:imageID", h.GetImage)
	v1.POST("/image", h.CreateImage)
	v1.DELETE("/image/:imageID", h.DeleteImage)
}
