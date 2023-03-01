package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pouyam79i/Cloud_Computing_HW/internal/app/server/config"
)

// This is a simple test handler!
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World from MY SERVER!!!!")
}

// TODO: complete sign in - make user login!
func SignIn(c echo.Context) error {
	fmt.Println("Received a POST data in sign in, Detail:")
	data := &config.SingInInfo{}
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		resErr := &config.ClientMSG{
			Type: "Respond Parsing Error",
			Info: "Error While Parsing Json File:\n" + err.Error(),
		}
		fmt.Println("Received MSG with Error:\n", resErr)
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		convertErr := c.JSON(http.StatusBadRequest, resErr)	
		if convertErr != nil{
			fmt.Println("Error in responding with json:\n", convertErr.Error())
		}
		return convertErr
	} else {
		fmt.Println("Received Data", "\nEmail: ", data.Email, "\nPassword: ", data.Password)
	}
	resSuc := &config.ClientMSG{
		Type: "Data Received",
		Info: "MSG received and parsed successfully!",
	}
	return c.JSON(http.StatusOK, resSuc)
}
