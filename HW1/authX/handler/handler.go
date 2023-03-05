package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pouyam79i/Cloud_Computing_HW/authX/api"
	"github.com/pouyam79i/Cloud_Computing_HW/authX/config"
)

func SingIn(c echo.Context) error {

	receivedKey := c.QueryParam("key")

	configs, err := config.GetConfigs()

	if err != nil {
		res := config.API_SIGNIN_RESULT{
			Result: false,
			Token:  "",
			Info:   "Cannot Load Key!. reason: " + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	mainKey := configs.ApiConnections.MainKey

	if strings.Compare(mainKey, receivedKey) != 0 {
		res := config.API_SIGNIN_RESULT{
			Result: false,
			Token:  "",
			Info:   "false api key",
		}
		return c.JSON(http.StatusUnauthorized, res)
	}

	reqBody := config.API_SIGNIN_REQUEST{}
	err = c.Bind(&reqBody)
	if err != nil {
		res := config.API_SIGNIN_RESULT{
			Result: false,
			Token:  "",
			Info:   "Decoding JSON failed!. reason: " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	token, err := api.SingInUser(reqBody)
	res := config.API_SIGNIN_RESULT{}
	code := http.StatusOK
	if err != nil {
		res.Result = false
		res.Token = ""
		res.Info = err.Error()
		code = http.StatusForbidden
	} else {
		res.Result = true
		res.Token = token
		res.Info = "successful"
	}
	return c.JSON(code, res)
}
