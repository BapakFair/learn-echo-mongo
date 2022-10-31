package query

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"learn-echo-mongo/config"
	"learn-echo-mongo/handlers"
	"log"
	"net/http"
)

var ctx = context.Background()

func CreateProduct(c echo.Context) error {
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var reqBody handlers.Product
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	fmt.Println("ini request isine: ", reqBody)
	data, err := db.Collection("products").InsertOne(ctx, handlers.Product{
		Name:        reqBody.Name,
		Price:       reqBody.Price,
		Currency:    reqBody.Currency,
		Discount:    reqBody.Discount,
		Vendor:      reqBody.Vendor,
		Accessories: reqBody.Accessories,
		IsEssential: false,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
	return c.JSON(http.StatusOK, data)
}
