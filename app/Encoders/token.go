package Encoders

import (
"io"
"github.com/Abdinour/learnGo/app/models"
"io/ioutil"
"encoding/json"
	"log"
)

func EncoderToken(body io.ReadCloser) (token models.Token) {
var data, _ =ioutil.ReadAll(body)

if err := json.Unmarshal(data, &token); err!=nil{
	log.Println(err)
return
}
return
}

