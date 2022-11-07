package controllers

import (
	"github.com/labstack/echo/v4"
	"learn-echo-mongo/models"
	"learn-echo-mongo/query"
	"log"
	"net/http"
)

func CreateProduct(c echo.Context) error {
	var reqBody models.Product
	if err := c.Bind(&reqBody); err != nil {
		log.Fatal(err.Error())
	}
	result, err := query.CreateProduct(reqBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func GetAllProducts(c echo.Context) error {
	result, err := query.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	c.Request()
	return c.JSON(http.StatusOK, result)
}

func GetProductById(c echo.Context) error {
	id := c.QueryParam("id")
	result, err := query.GetProductById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	c.Request()
	return c.JSON(http.StatusOK, result)
}
