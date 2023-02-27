package model

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func BuildServer(debug bool) *echo.Echo {
	fmt.Println("Building Server ...")
	e := echo.New()
	if debug {
		AttachALL(e)
	} else {
		AttachMain(e)
	}
	fmt.Println("Done Building!")
	return e
}
