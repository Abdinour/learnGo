package Encoders

import (
	"io"
	"github.com/Abdinour/learnGo/app/models"
	"io/ioutil"
	"encoding/json"
)

func EncoderSingleUsers(body io.ReadCloser) (user models.User) {
	var data, _ =ioutil.ReadAll(body)

	if err := json.Unmarshal(data, &user); err!=nil{
		return
	}
	return user
}
