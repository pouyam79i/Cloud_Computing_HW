package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pouyam79i/Cloud_Computing_HW/main/HW2/step2/code/api"
	"github.com/pouyam79i/Cloud_Computing_HW/main/HW2/step2/code/config"
)

// Call on rebrandly api and return result
// 1 - check if it exists in 'redis' and return if true.
// 2 - else retrieve data from rebrandly.com, also save it in cache redis.
func CallRebrandly(c echo.Context) error {

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unable to Get Hostname"
		err = nil
	}

	userReq := config.RequestAPI{}
	err = json.NewDecoder(c.Request().Body).Decode(&userReq)

	if err != nil {
		response := &config.ResponseAPI{
			LongURL:  "",
			ShortURL: "",
			IsCached: false,
			Hostname: hostname,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	var isCached bool = true
	shortenedUrl, err := api.GetRedis(userReq.URL)
	if err != nil {
		fmt.Println("failed to get data from redis. err: ", err.Error())
		shortenedUrl, err = api.CallRebrandlyAPI(userReq.URL)
		isCached = false
	}

	if err != nil {
		response := &config.ResponseAPI{
			LongURL:  userReq.URL,
			ShortURL: err.Error(),
			IsCached: false,
			Hostname: hostname,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if !isCached {
		err = api.SendRedis(userReq.URL, shortenedUrl)
		if err != nil {
			fmt.Println("failed to save data at redis. err: ", err.Error())
		}
	}

	response := &config.ResponseAPI{
		LongURL:  userReq.URL,
		ShortURL: shortenedUrl,
		IsCached: isCached,
		Hostname: hostname,
	}

	return c.JSON(http.StatusOK, response)
}
