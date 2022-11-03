package routes

import (
	"github.com/labstack/echo/v4"
	"learn-echo-mongo/controllers"
	"learn-echo-mongo/query"
	"net/http"
)

func Init() *echo.Echo {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hallo this is echo")
	})
	e.GET("/api/products", controllers.GetAllProducts)
	e.GET("/api/products/:id", query.GetProducts)
	e.POST("/api/products", controllers.CreateProduct)

	return e
}
