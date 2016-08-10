package Interceptors

import (
	"github.com/revel/revel"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/Abdinour/learnGo/app/util"
	"strconv"
	"time"
)

type JWTAuthorization struct {

   *revel.Controller
}

func (c JWTAuthorization) CheckUser() revel.Result {
	var tokenstring = c.Request.Header.Get("token")
	token, err := jwt.Parse(tokenstring, func(token  *jwt.Token) (interface{}, error) {
		//Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method:%v", token.Header["alg"])

		}
		appSecret, _ := revel.Config.String("app.secret");
		return []byte(appSecret), nil
	})
	if err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var expDate = time.Unix(int64(claims["exp"].(float64)) , 0)
			if expDate.Before(time.Now()) {
				return c.RenderJson(util.ResponseError("expired Token"))
			}
			c.Session["id"] = strconv.Itoa(int(claims["id"].(float64)));
			c.Session["email"] = claims["email"].(string);
			return nil
		}
		return c.RenderJson(util.ResponseError("Invalid token key"))

	}else {
		return c.RenderJson(util.ResponseError("Not Found token key"))
	}
}


func init()  {
	revel.InterceptMethod(JWTAuthorization.CheckUser,revel.BEFORE);

}
