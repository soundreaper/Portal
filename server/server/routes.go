package server

import (
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
	s.e.GET("/hello", h.Hello)

	// Group Under API V1 in case we want to change in the future
	v1 := s.e.Group("/api/v1")

	// Enables auth middleware so that all routes below require authorization
	mw := getJwtMiddleware()
	v1.Use(echo.WrapMiddleware(mw.Handler))

	// User routes
	//v1.GET("/user", h.GetUser)
	//v1.POST("/user", h.CreateUser)
	//v1.DELETE("/user", h.DeleteUser)

	// Image routes
	v1.GET("/images", h.Upload)
	v1.POST("/upload", h.Images)
	v1.DELETE("/image/:objectID", h.Delete)
}
