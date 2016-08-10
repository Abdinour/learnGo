package Encoders

import (
	"io"
	"github.com/Abdinour/learnGo/app/models"
	"io/ioutil"
	"encoding/json"
)

func EncoderSinglePost(body io.ReadCloser) (post models.Post) {
	var data, _ =ioutil.ReadAll(body)

	if err := json.Unmarshal(data, &post); err!=nil{
		return
	}
	return post
}
