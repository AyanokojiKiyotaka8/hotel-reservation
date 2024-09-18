package main

import (
	"context"
	"flag"
	"log"

	"github.com/AyanokojiKiyotaka8/hotel-reservation/api"
	"github.com/AyanokojiKiyotaka8/hotel-reservation/db"
	"github.com/AyanokojiKiyotaka8/hotel-reservation/api/middleware"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddress := flag.String("listenAddress", ":3000", "The listen address of API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	// storages
	userStore := db.NewMongoUserStore(client, db.DBNAME)
	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME, hotelStore)
	store := &db.Store{
		Hotel: hotelStore,
		Room: roomStore,
		User: userStore,
	}

	app := fiber.New(config)
	auth := app.Group("/api")
	apiv1 := app.Group("/api/v1", middleware.JWTAuthentication(userStore))

	// handlers
	userHandler := api.NewUserHandler(userStore)
	hotelHandler := api.NewHotelHandler(store)
	authHandler := api.NewAuthHandler(userStore)
	roomHandler := api.NewRoomHandler(store)

	// auth api
	auth.Post("/auth", authHandler.HandleAuthenticate)

	// user api
	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	// hotel api
	apiv1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	apiv1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)

	// room api
	apiv1.Post("/room/:id/book", roomHandler.HandleBookRoom)

	app.Listen(*listenAddress)
}