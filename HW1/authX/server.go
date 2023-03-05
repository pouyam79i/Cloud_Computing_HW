package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pouyam79i/Cloud_Computing_HW/authX/config"
	"github.com/pouyam79i/Cloud_Computing_HW/authX/db"
	"github.com/pouyam79i/Cloud_Computing_HW/authX/handler"
)

func main() {

	fmt.Println("Booting AuthX Server ...")

	configs, err := config.LoadAll()
	if err != nil {
		panic(err.Error())
	}

	_ = db.Connect(configs.DBI)

	e := echo.New()

	// attaching api handlers
	e.POST("/signin", handler.SingIn)

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	var address string = configs.Server.IP + ":" + configs.Server.Port

	// fmt.Println("Main Connection Key: ", configs.ApiConnections.MainKey)

	e.Logger.Fatal(e.Start(address))
}
