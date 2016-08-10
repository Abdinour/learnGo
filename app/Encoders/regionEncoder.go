package Encoders


import (
	"io"
	"github.com/Abdinour/learnGo/app/models"
	"io/ioutil"
	"encoding/json"
)

func EncodeRegion(body io.ReadCloser) (region models.Region) {
	var data, _ =ioutil.ReadAll(body)

	if err := json.Unmarshal(data,&region); err!=nil{
		return
	}
	return
}
