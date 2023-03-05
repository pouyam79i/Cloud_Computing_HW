package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pouyam79i/Cloud_Computing_HW/server/internal/app/server/config"
	"github.com/pouyam79i/Cloud_Computing_HW/server/internal/app/server/handler/api"
)

// This is a simple test handler!
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World from MY SERVER!!!!")
}

// TODO: complete sign in - make user login!
func SignIn(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	data := config.SingInInfo{}
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		resErr := &config.ClientMSG{
			Result: false,
			Token:  "",
			Info:   "Error While Parsing Json File:\n" + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, resErr)
	} else {
		fmt.Println("Received Data", "\nEmail: ", data.Email, "\nPassword: ", data.Password)
	}

	// TODO: connect to authx
	token, err := api.Authx_SingIn(data)
	res := config.ClientMSG{}
	code := http.StatusOK
	if err != nil {
		res.Result = false
		res.Token = ""
		res.Info = err.Error()
		code = http.StatusForbidden
	} else {
		res.Result = true
		res.Token = token
		res.Info = "Successful"
	}

	return c.JSON(code, res)
}

func TestValidator(c echo.Context) error {
	data := config.JustToken{}
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		resErr := &config.ClientMSG{
			Result: false,
			Token:  "",
			Info:   "Error While Parsing Json File:\n" + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, resErr)
	} else {
		fmt.Println("Received Data from user", "\nGiven Token: ", data.Token)
	}

	isValid, err := api.Authx_Validate(data.Token)
	info := "successful"

	if err != nil {
		info = "Err: " + err.Error()
	}

	resErr := &config.ClientMSG{
		Result: isValid,
		Token:  "true means given token was valid",
		Info:   info,
	}
	return c.JSON(http.StatusOK, resErr)
}
