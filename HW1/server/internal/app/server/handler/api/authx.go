package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/pouyam79i/Cloud_Computing_HW/server/internal/app/server/config"
)

// his is a randomly generate hast
// You should put same key as authx main key here!

func Authx_SingIn(userInfo config.SingInInfo) (string, error) {

	ui_str, err := json.Marshal(userInfo)
	if err != nil {
		return "", errors.New("cannot marshal given json struct")
	}

	payload := bytes.NewBuffer(ui_str)
	client := http.Client{}

	req, err := http.NewRequest(http.MethodPost, config.API_AUTHX_SIGNIN+config.AUTHX_KEY, payload)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("Failed to build req. reason:\n", err.Error())
		return "", err
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Failed to read respond from Authx. Reason:\n", err.Error())
	}

	byteString, _ := io.ReadAll(res.Body)

	var r config.ClientMSG
	err = json.Unmarshal(byteString, &r)

	if err != nil {
		return "", errors.New("cannot Unmarshal given json string")
	}

	if r.Result {
		return r.Token, nil
	} else {
		return "", errors.New(r.Info)
	}

}

// Validate user token
func Authx_Validate(token string) (bool, error) {
	return false, nil
}
