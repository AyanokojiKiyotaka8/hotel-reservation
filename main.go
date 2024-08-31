package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/AyanokojiKiyotaka8/hotel-reservation/api"
	"github.com/AyanokojiKiyotaka8/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)

	user := types.User{
		FirstName: "Makha",
		LastName: "Zadrbek",
	}

	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	var makha types.User
	if err := coll.FindOne(ctx, bson.M{}).Decode(&makha); err != nil {
		log.Fatal(err)
	}
	fmt.Println(makha)


	listenAddress := flag.String("listenAddress", ":3000", "The listen address of API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	fmt.Println("Starting at port:3000")
	app.Listen(*listenAddress)
}