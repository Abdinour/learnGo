package Encoders
import (
	"io"
	"github.com/Abdinour/learnGo/app/models"
	"io/ioutil"
	"encoding/json"
)

func LikeEncoder(body io.ReadCloser) (like models.Likes) {
	var data, _ =ioutil.ReadAll(body)

	if err := json.Unmarshal(data,&like); err!=nil{
		return
	}
	return
}