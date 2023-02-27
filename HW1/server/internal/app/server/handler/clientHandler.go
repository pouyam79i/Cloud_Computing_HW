package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// This is a simple test handler!
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World from MY SERVER!!!!")
}
