package Encoders


import (
	"io"
	"github.com/Abdinour/learnGo/app/models"
	"io/ioutil"
	"encoding/json"
)

func EncodeTax(body io.ReadCloser) (tax models.Tax) {
	var data, _ =ioutil.ReadAll(body)

	if err := json.Unmarshal(data,&tax); err!=nil{
		return
	}
	return
}
