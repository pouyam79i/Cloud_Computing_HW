package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pouyam79i/Cloud_Computing_HW/internal/app/server/config"
	"github.com/pouyam79i/Cloud_Computing_HW/internal/app/server/handler/api"
)

// TODO: complete code uploading mechanism
func UploadCode(c echo.Context) error {
	reqBody := config.ClientCode{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := config.ClientMSG{
			Type: "Decoding JSON failed!",
			Info: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	api.SendCodeX(reqBody)
	res := config.ClientMSG{
		Type: "Code Received",
		Info: "We are sending your code to codex api. Please be patient, the result of your code will be sent to your email",
	}
	return c.JSON(http.StatusOK, res)
}
