package controllers

import (

	"github.com/revel/revel"
	"github.com/Abdinour/learnGo/app/models"
	"github.com/Abdinour/learnGo/app"
	"github.com/Abdinour/learnGo/app/util"
	"github.com/Abdinour/learnGo/app/Encoders"
	"log"
	"strconv"
)

type Postcontroller struct {

	//Interceptors.JWTAuthorization
	*revel.Controller
}

func (c Postcontroller)  Index() revel.Result{
	var posts []models.Post
	var limitQuery =c.Request.URL.Query().Get("limit");
	if limitQuery =="" {
		limitQuery= "0"
	}
	var offsetQuery =c.Request.URL.Query().Get("offset");
	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&posts).RowsAffected; founded < 1 {
              return c.RenderJson( util.ResponseError("No founded posts"))
	}
	 for i, post := range posts {
		app.Db.First(&posts[i].User, post.UserID)
		 posts[i].User.Password =""
	}
       return c.RenderJson(posts);
}

func (c Postcontroller) Create() revel.Result  {
	var post = Encoders.EncoderSinglePost(c.Request.Body);
	if post.Body =="" &&post.Title==""{
		return c.RenderJson(util.ResponseError("Post Information not founded"));
	}
	post.UserID, _ = strconv.ParseInt(c.Session["id"],10, 0)
	if err :=app.Db.Create(&post).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("Post Creation failed"));

	}
	return  c.RenderJson(util.ResponseSuccess(post))
}

func (c Postcontroller) Update() revel.Result  {
var update = Encoders.EncoderSinglePost(c.Request.Body);
	var id int
	var post models.Post
	// bind params
	c.Params.Bind(&id, "id")
	if rowscount := app.Db.First(&post, id).RowsAffected; rowscount < 1 {
		return c.RenderJson(util.ResponseError("Post Information not founded"));
	}
	if err := app.Db.Model(&post).Update(&update).Error; err != nil {
		log.Println(err);
	return c.RenderJson(util.ResponseError("Post Update failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(post))
}

func (c Postcontroller) Delete() revel.Result {
	var(
	id int
	 post models.Post
	)
	c.Params.Bind(&id,"id");
	if rowscount := app.Db.First(&post, id).RowsAffected; rowscount < 1 {
		return c.RenderJson(util.ResponseError("Post Information not founded"));
	}
             // bind

	if err := app.Db.Delete(&post).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Post Delete failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(post))
}




