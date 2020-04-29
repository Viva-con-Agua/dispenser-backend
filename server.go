package main

import (
	controller "dispenser-backend/controllers"
	"dispenser-backend/database"
	"dispenser-backend/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {

	// intial loading function
	utils.LoadConfig()
	database.ConnectMongo()
	//create echo server
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	apiV1 := e.Group("/api/v1")
	apiV1.POST("/navigation", controller.NavigationInsert)
	apiV1.GET("/navigation/:name", controller.NavigationGetByName)
	e.Logger.Fatal(e.Start(":1323"))
}
