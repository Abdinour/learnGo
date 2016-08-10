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

type Regioncontroller struct {
	*revel.Controller
}

func (c Regioncontroller)  Index() revel.Result{
	var region []models.Region
	var limitQuery =c.Request.URL.Query().Get("limit");
	if limitQuery =="" {
		limitQuery= "0"
	}
	var offsetQuery =c.Request.URL.Query().Get("offset");
	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&region).RowsAffected; founded < 1 {
		return c.RenderJson( util.ResponseError("No founded Comment"))
	}

	return c.RenderJson(util.ResponseSuccess(region));
}

func (c Regioncontroller) Create() revel.Result  {
	var region = Encoders.EncodeRegion(c.Request.Body);
	if region.Longtitude =="" &&region.Latitude==""{
		return c.RenderJson(util.ResponseError("Post Information not founded"));
	}
	region.CitizenID, _ = strconv.ParseInt(c.Session["id"],10, 0)
	if err :=app.Db.Create(&region).Error; err != nil {
		return c.RenderJson(util.ResponseError("Citizen Creation failed"));

	}
	return  c.RenderJson(util.ResponseSuccess(region))
}

func (c Regioncontroller) Update() revel.Result  {
	var update = Encoders.EncodeRegion(c.Request.Body);
	var id int
	var region models.Region
	// bind params
	c.Params.Bind(&id, "id")
	if rowscount := app.Db.First(&region, id).RowsAffected; rowscount < 1 {
		return c.RenderJson(util.ResponseError("Comment Information not founded"));
	}
	if err := app.Db.Model(&region).Update(&update).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Comment Update failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(region))
}

func (c Regioncontroller) Delete() revel.Result {
	var(
		id int
		region models.Region
	)
	c.Params.Bind(&id,"id");
	if rowscount := app.Db.First(&region, id).RowsAffected; rowscount < 1 {
		return c.RenderJson(util.ResponseError("Comment Information not founded"));
	}
	// bind

	if err := app.Db.Delete(&region).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Comment Delete failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(region))
}




