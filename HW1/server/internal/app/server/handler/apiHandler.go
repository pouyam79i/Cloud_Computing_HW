package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pouyam79i/Cloud_Computing_HW/server/internal/app/server/config"
	"github.com/pouyam79i/Cloud_Computing_HW/server/internal/app/server/handler/api"
)

// TODO: complete code uploading mechanism
func UploadCode(c echo.Context) error {
	reqBody := config.ClientCode{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := config.ClientMSG{
			Result: false,
			Info:   err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	// TODO: save the file and tell the job scheduler!
	// api.SendCodeX(reqBody)
	api.SendDataToJobBuilder(reqBody)
	res := config.ClientMSG{
		Result: true,
		Info:   "We are sending your code to codex api. Please be patient, the result of your code will be sent to your email",
	}
	return c.JSON(http.StatusOK, res)
}
