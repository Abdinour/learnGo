package app

import (
	"github.com/revel/revel"
	"github.com/jinzhu/gorm"
	"github.com/Abdinour/learnGo/app/models"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"log"

)
var Db *gorm.DB

func init() {
	//Filters is the defaukty set of global filters
	revel.Filters = [] revel.Filter{
		revel.PanicFilter,
		revel.RouterFilter,
		revel.FilterConfiguringFilter,
		revel.ParamsFilter,
		revel.SessionFilter,
		revel.FlashFilter,
		revel.ValidationFilter,
		revel.I18nFilter,
		HeaderFilter,
		revel.InterceptorFilter,
		revel.CompressFilter,
		revel.ActionInvoker,

	}

	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
	revel.OnAppStart(func() {
		var err error
		Db, err = gorm.Open("sqlite3", "./src/github.com/Abdinour/learnGo/Database/storage.db");
		if err != nil {
			log.Fatal(err);
		}
		Db.CreateTable(&models.User{});
		Db.CreateTable(&models.Citizen{});
		Db.CreateTable(&models.Property{});
		Db.CreateTable(&models.Region{});
		Db.CreateTable(&models.Tax{});
		Db.DB().SetMaxIdleConns(10);
		Db.DB().SetMaxOpenConns(100);
		//Disable Table name's Pluralization
		Db.SingularTable(false)
		Db.LogMode(true)


	})
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
