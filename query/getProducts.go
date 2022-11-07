package query

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"learn-echo-mongo/config"
	"learn-echo-mongo/models"
	"log"
	"net/http"
	"time"
)

func GetProductById(id string) (models.Response, error) {
	var res models.Response
	var product models.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	if err := db.Collection("products").FindOne(ctx, bson.M{"_id": objId}).Decode(&product); err != nil {
		log.Fatal(err)
	}

	res.Status = http.StatusOK
	res.Message = "Get data success"
	res.Data = product

	return res, nil
}

func GetAllProducts() (models.Response, error) {
	var res models.Response
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := config.Connect()
	if err != nil {
		return res, err
	}

	defer cancel()

	data, err := db.Collection("products").Find(ctx, bson.M{})
	if err != nil {
		return res, err
	}
	var dataFinal []bson.M
	if err = data.All(ctx, &dataFinal); err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Get data success"
	res.Data = dataFinal

	return res, nil
}
