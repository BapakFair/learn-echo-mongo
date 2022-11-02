package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"learn-echo-mongo/config"
	"learn-echo-mongo/query"
	"log"
	"net/http"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func main() {
	_, err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasss...")

	e := echo.New()
	e.GET("/api/product", query.GetAllProducts)
	e.GET("/api/product/:id", query.GetProducts)
	e.POST("/api/products", query.CreateProduct)

	e.Start(":8989")
}
