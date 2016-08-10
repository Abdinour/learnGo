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

type Citizencontroller struct {

	*revel.Controller
}

func (c Citizencontroller)  Index() revel.Result{
	var citizens []models.Citizen
	var limitQuery =c.Request.URL.Query().Get("limit");
	if limitQuery =="" {
		limitQuery= "0"
	}
	var offsetQuery =c.Request.URL.Query().Get("offset");
	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&citizens).RowsAffected; founded < 1 {
		return c.RenderJson( util.ResponseError("No founded Comment"))
	}
	return c.RenderJson(util.ResponseSuccess(citizens));
}

func (c Citizencontroller) Create() revel.Result  {
	var citizen = Encoders.EncodeCitizen(c.Request.Body);
	citizen.UserID, _ = strconv.ParseInt(c.Session["id"],10, 0)
	if err :=app.Db.Create(&citizen).Error; err != nil {
		return c.RenderJson(util.ResponseError("Citizen Creation failed"));

	}
	return  c.RenderJson(util.ResponseSuccess(citizen))
}

func (c Citizencontroller) Update() revel.Result  {
	var update = Encoders.EncodeCitizen(c.Request.Body);
	var id int
	var citizen models.Citizen
	// bind params
	c.Params.Bind(&id, "id")
	if rowscount := app.Db.First(&citizen, id).RowsAffected; rowscount < 1 {
		return c.RenderJson(util.ResponseError("Comment Information not founded"));
	}
	if err := app.Db.Model(&citizen).Update(&update).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Comment Update failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(citizen))
}

func (c Citizencontroller) Delete() revel.Result {
	var(
		id int
		citizen models.Citizen
	)
	c.Params.Bind(&id,"id");
	if rowscount := app.Db.First(&citizen, id).RowsAffected; rowscount < 1 {
		return c.RenderJson(util.ResponseError("Comment Information not founded"));
	}
	// bind

	if err := app.Db.Delete(&citizen).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Comment Delete failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(citizen))
}




