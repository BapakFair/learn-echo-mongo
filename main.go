package main

import (
	"github.com/labstack/echo/v4"
	"learn-echo-mongo/config"
	"learn-echo-mongo/routes"
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

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8989"))
}
