package server

import (
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

	// Setup our route handlers
	h := handlers.NewHandler(s.db)

	// Hello route for a status check
	s.e.GET("/hello", h.Hello)

	// Auth routes
	s.e.GET("/auth/:provider/callback", h.GetCallback)
	s.e.GET("/logout/:provider", h.Logout)
	s.e.GET("/auth/:provider", h.Login)

	// Group Under API V1 in case we want to change in the future
	v1 := s.e.Group("/api/v1")

	// Image routes
	v1.GET("/images", h.Upload)
	v1.POST("/upload", h.Images)
	v1.DELETE("/image/:objectID", h.Delete)
}
