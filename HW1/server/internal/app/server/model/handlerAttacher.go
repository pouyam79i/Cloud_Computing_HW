package model

import (
	"github.com/labstack/echo/v4"
	"github.com/pouyam79i/Cloud_Computing_HW/server/internal/app/server/handler"
)

// TODO: add all routes to your echo server here
// attacking all routes here
// include everything for debugging
func AttachALL(e *echo.Echo) {
	e.GET("/hello", handler.HelloWorld)
	e.POST("/signin", handler.SignIn)
	e.POST("/codex", handler.UploadCode)
}

// TODO: attacking your business' routes here
func AttachMain(e *echo.Echo) {
	e.POST("/signin", handler.SignIn)
	e.POST("/codex", handler.UploadCode)
}
