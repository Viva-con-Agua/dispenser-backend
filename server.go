package main

import (
	controller "dispenser-backend/controllers"
	"dispenser-backend/database"
	"dispenser-backend/utils"

	"github.com/Viva-con-Agua/echo-pool/auth"
	"github.com/Viva-con-Agua/echo-pool/config"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	config.LoadConfig()
	database.ConnectMongo()
	//create echo server
	store := auth.RedisSession()
	//create echo server
	e := echo.New()
	m := middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     utils.Config.Alloworigins,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	})
	e.Use(m)
	e.Use(store)
	e.Validator = &CustomValidator{validator: validator.New()}
	apiV1 := e.Group("/v1")
	apiV1.POST("/navigation", controller.NavigationInsert)
	apiV1.PUT("/navigation", controller.NavigationUpdate)
	apiV1.GET("/navigation/:name", controller.NavigationGetByName)
	e.Logger.Fatal(e.Start(":1323"))
}
