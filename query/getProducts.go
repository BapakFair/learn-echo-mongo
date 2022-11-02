package query

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"learn-echo-mongo/config"
	"learn-echo-mongo/handlers"
	"log"
	"net/http"
	"time"
)

func GetProducts(c echo.Context) error {
	fmt.Println("manuk")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	productId := c.Param("id")
	var product handlers.Product
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(productId)
	fmt.Println("ini objid: ", product)

	if err = db.Collection("products").FindOne(ctx, bson.M{"_id": objId}).Decode(&product); err != nil {
		log.Fatal(err)
	}

	fmt.Println(product)
	return c.JSON(http.StatusOK, &echo.Map{"data": product})
}

func GetAllProducts(c echo.Context) error {
	fmt.Println("manuk")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var product handlers.Product
	defer cancel()

	fmt.Println("ini objid: ", product)

	data, err := db.Collection("products").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var dataFinal []bson.M
	if err = data.All(ctx, &dataFinal); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, &echo.Map{"data": dataFinal})
}
