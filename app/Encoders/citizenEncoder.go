package Encoders


import (
	"io"
	"github.com/Abdinour/learnGo/app/models"
	"io/ioutil"
	"encoding/json"
)

func EncodeCitizen(body io.ReadCloser) (citizen models.Citizen) {
	var data, _ =ioutil.ReadAll(body)

	if err := json.Unmarshal(data,&citizen); err!=nil{
		return
	}
	return
}
