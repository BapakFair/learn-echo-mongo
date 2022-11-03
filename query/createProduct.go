package query

import (
	"context"
	"github.com/labstack/echo/v4"
	"learn-echo-mongo/config"
	"learn-echo-mongo/models"
	"log"
	"net/http"
)

var ctx = context.Background()

func CreateProduct(c echo.Context) (models.Response, error) {
	var res models.Response

	db, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var reqBody models.Product
	if err := c.Bind(&reqBody); err != nil {
		log.Fatal(err.Error())
	}

	data, err := db.Collection("products").InsertOne(ctx, models.Product{
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

	res.Status = http.StatusOK
	res.Message = "Insert data success"
	res.Data = data

	return res, nil
}
