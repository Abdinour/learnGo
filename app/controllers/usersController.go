package controllers

import (
	"github.com/revel/revel"
	"log"
	"github.com/Abdinour/learnGo/app"
	"github.com/Abdinour/learnGo/app/Encoders"
	"github.com/Abdinour/learnGo/app/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/Abdinour/learnGo/app/util"
	"time"
)

type UsersController struct {

	*revel.Controller
}

func (c UsersController)  Create() revel.Result {
	var user = Encoders.EncoderSingleUsers(c.Request.Body)
	if user.Email =="" || user.Password ==""{
		return c.RenderJson(util.ResponseError("User Creation is empty"));
	}
	if err :=app.Db.Create(&user).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("User Creation failed"));

	}
	return c.RenderJson(util.ResponseSuccess(user))
}
func (c UsersController) Login() revel.Result{
	var user = Encoders.EncoderSingleUsers(c.Request.Body);
	if user.Email =="" || user.Password ==""{
		return c.RenderJson(util.ResponseError("User Creation is empty"));
	}
	if founded := app.Db.Where(&user).First(&user).RowsAffected; founded  < 1 {
		return c.RenderJson(util.ResponseError("User Not Founded"));

	}

	token :=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"id":  user.ID,
		"email": user.Email,
		"exp":time.Now().Add(time.Hour * 24).Unix(),
	})
        appSecret,_ := revel.Config.String("app.secret");
	tokenString, err := token.SignedString([]byte(appSecret));
	if err!= nil{
		return c.RenderJson(util.ResponseError("Key Generation Failed"));
	}
  var tokenmodel models.Token
	tokenmodel.Name= user.Name
	tokenmodel.Email  =user.Email
	tokenmodel.Token   = tokenString
	return c.RenderJson(util.ResponseSuccess(tokenmodel))


}
func (c UsersController)  Index() revel.Result{
	var users []models.User
	var limitQuery =c.Request.URL.Query().Get("limit");
	if limitQuery =="" {
		limitQuery= "0"
	}
	var offsetQuery =c.Request.URL.Query().Get("offset");
	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&users).RowsAffected; founded < 1 {
		return c.RenderJson( util.ResponseError("No founded Comment"))
	}
	return c.RenderJson(util.ResponseSuccess(users));

}

func (c UsersController) Delete() revel.Result {
	var(
		id int
		user models.User
	)
	c.Params.Bind(&id,"id");
	if rowscount := app.Db.First(&user, id).RowsAffected; rowscount < 1 {
		return c.RenderJson(util.ResponseError("User Information not founded"));
	}
	// bind

	if err := app.Db.Delete(&user).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("User Delete failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(user))
}

func (c UsersController) Update() revel.Result  {
	var update = Encoders.EncoderSingleUsers(c.Request.Body);
	var id int
	var user models.User
	// bind params
	c.Params.Bind(&id, "id")
	if rowscount := app.Db.First(&user, id).RowsAffected; rowscount < 1 {
		return c.RenderJson(util.ResponseError("Comment Information not founded"));
	}
	if err := app.Db.Model(&user).Update(&update).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Comment Update failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(user))
}







