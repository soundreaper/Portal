package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Hello(c echo.Context) error {
	// hello world :)
	res := map[string]string{"msg": "hello world :)"}

	return c.JSON(http.StatusOK, res)
}
