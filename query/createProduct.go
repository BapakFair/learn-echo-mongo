package query

import (
	"context"
	"learn-echo-mongo/config"
	"learn-echo-mongo/models"
	"log"
	"net/http"
	"time"
)

func CreateProduct(reqBody models.Product) (models.Response, error) {
	var res models.Response
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	db, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cancel()

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

	res.Status = http.StatusCreated
	res.Message = "Insert data success"
	res.Data = data

	return res, nil
}
