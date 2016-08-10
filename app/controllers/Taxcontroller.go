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

type Taxcontroller struct {

	*revel.Controller
}

func (c Taxcontroller)  Index() revel.Result{
	var tax []models.Tax
	var limitQuery =c.Request.URL.Query().Get("limit");
	if limitQuery =="" {
		limitQuery= "0"
	}
	var offsetQuery =c.Request.URL.Query().Get("offset");
	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&tax).RowsAffected; founded < 1 {
		return c.RenderJson( util.ResponseError("No founded Comment"))
	}

	return c.RenderJson(util.ResponseSuccess(tax));
}

func (c Taxcontroller) Create() revel.Result  {
	var tax = Encoders.EncodeTax(c.Request.Body);
	tax.PropertyID, _ = strconv.ParseInt(c.Session["id"],10, 0)
	if err :=app.Db.Create(&tax).Error; err != nil {
		return c.RenderJson(util.ResponseError("Tax Creation failed"));

	}
	return  c.RenderJson(util.ResponseSuccess(tax))
}

func (c Taxcontroller) Update() revel.Result  {
	var update = Encoders.EncodeTax(c.Request.Body);
	var id int
	var tax models.Tax
	// bind params
	c.Params.Bind(&id, "id")
	if rowscount := app.Db.First(&tax, id).RowsAffected; rowscount < 1 {
		return c.RenderJson(util.ResponseError("Comment Information not founded"));
	}
	if err := app.Db.Model(&tax).Update(&update).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Comment Update failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(tax))
}

func (c Taxcontroller) Delete() revel.Result {
	var(
		id int
		tax models.Tax
	)
	c.Params.Bind(&id,"id");
	if rowscount := app.Db.First(&tax, id).RowsAffected; rowscount < 1 {
		return c.RenderJson(util.ResponseError("Comment Information not founded"));
	}
	// bind

	if err := app.Db.Delete(&tax).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Comment Delete failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(tax))
}




