
package Encoders


import (
"io"
"github.com/Abdinour/learnGo/app/models"
"io/ioutil"
"encoding/json"
)

func EncodeProperty(body io.ReadCloser) (property models.Property) {
	var data, _ =ioutil.ReadAll(body)

	if err := json.Unmarshal(data,&property); err!=nil{
		return
	}
	return
}

