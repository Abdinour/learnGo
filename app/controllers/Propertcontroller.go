
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

type Propertycontroller struct {
	*revel.Controller
}

func (c Propertycontroller)  Index() revel.Result{
	var properties []models.Property
	var limitQuery =c.Request.URL.Query().Get("limit");
	if limitQuery =="" {
		limitQuery= "0"
	}
	var offsetQuery =c.Request.URL.Query().Get("offset");
	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&properties).RowsAffected; founded < 1 {
		return c.RenderJson( util.ResponseError("No founded Comment"))
	}

	return c.RenderJson(util.ResponseSuccess(properties));
}

func (c Propertycontroller) Create() revel.Result  {
	var property = Encoders.EncodeProperty(c.Request.Body);
	property.TaxID, _ =strconv.ParseInt(c.Session["id"], 10, 0)
	if err :=app.Db.Create(&property).Error; err != nil {
		return c.RenderJson(util.ResponseError("property Creation failed"));

	}
	return  c.RenderJson(util.ResponseSuccess(property))
}

func (c Propertycontroller) Update() revel.Result  {
	var update = Encoders.EncodeProperty(c.Request.Body);
	var id int
	var property models.Property
	// bind params
	c.Params.Bind(&id, "id")
	if rowscount := app.Db.First(&property, id).RowsAffected; rowscount < 1 {
		return c.RenderJson(util.ResponseError("Comment Information not founded"));
	}
	if err := app.Db.Model(&property).Update(&update).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Comment Update failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(property))
}

func (c Propertycontroller) Delete() revel.Result {
	var(
		id int
		property models.Property
	)
	c.Params.Bind(&id,"id");
	if rowscount := app.Db.First(&property, id).RowsAffected; rowscount < 1 {
		return c.RenderJson(util.ResponseError("Comment Information not founded"));
	}
	// bind

	if err := app.Db.Delete(&property).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Comment Delete failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(property))
}





