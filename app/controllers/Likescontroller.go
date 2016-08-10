package controllers

import (
	"github.com/revel/revel"
	"github.com/Abdinour/learnGo/app/Interceptors"
	"github.com/Abdinour/learnGo/app/models"
	"github.com/Abdinour/learnGo/app"
	"github.com/Abdinour/learnGo/app/util"
	"log"
	"strconv"
)

type Likescontroller struct {

	Interceptors.JWTAuthorization
	*revel.Controller
}

func (c Likescontroller)  Index() revel.Result{
	var like models.Likes
	like.UserID, _ = strconv.ParseInt(c.Session["id"],10, 0)
	c.Params.Bind(&like.PostID,"id")
	log.Println(like);
	var likes  int64
	app.Db.Model(&models.Likes{}).Where(&likes).Count(&likes)

	return c.RenderJson(likes);
}

func (c Likescontroller) Create() revel.Result  {
	var likes  models.Likes
	likes.UserID, _ = strconv.ParseInt(c.Session["id"],10, 0)
	c.Params.Bind(&likes.PostID, "id")
	app.Db.First(&likes)
	if app.Db.NewRecord(&likes){
		if err :=app.Db.Create(&likes).Error; err != nil {
			return c.RenderJson(util.ResponseError("likes Creation failed"));

		}
	}

	return  c.RenderJson(util.ResponseSuccess(likes))
}

func (c Likescontroller) Delete() revel.Result {
	var like  models.Likes
	like.UserID, _ = strconv.ParseInt(c.Session["id"],10, 0)
	c.Params.Bind(&like.PostID, "id")
	if rowscount := app.Db.First(&like, like).RowsAffected; rowscount < 1 {
		return c.RenderJson(util.ResponseError("likes Information not founded"));
	}
	// bind

	if err := app.Db.Delete(&like).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("likes Delete failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(like))
}





