package controller

import (
	"dispenser-backend/database"
	"dispenser-backend/models"
	"net/http"

	"github.com/Viva-con-Agua/echo-pool/resp"
	"github.com/labstack/echo"
)

func NavigationInsert(c echo.Context) (err error) {
	body := new(models.Navigation)
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err = database.NavigationInsert(body); err != nil {
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError)
	}
	return c.JSON(http.StatusCreated, resp.Created())
}

func NavigationUpdate(c echo.Context) (err error) {
	body := new(models.Navigation)
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err = database.NavigationUpdate(body); err != nil {
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError)
	}
	return c.JSON(http.StatusOK, resp.Updated(body.Name))
}

func NavigationGetByName(c echo.Context) (err error) {
	name := c.Param("name")
	response, err := database.NavigationGetByName(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError)
	}
	return c.JSON(http.StatusOK, response)
}
