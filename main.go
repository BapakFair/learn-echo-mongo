package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"learn-echo-mongo/config"
	"learn-echo-mongo/query"
	"log"
)

func main() {

	_, err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasss...")

	e := echo.New()
	e.POST("api/products", query.CreateProduct)

	e.Start(":8989")
}
