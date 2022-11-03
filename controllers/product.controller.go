package controllers

import (
	"github.com/labstack/echo/v4"
	"learn-echo-mongo/query"
	"net/http"
)

func CreateProduct(c echo.Context) error {
	result, err := query.CreateProduct(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllProducts(c echo.Context) error {
	result, err := query.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
