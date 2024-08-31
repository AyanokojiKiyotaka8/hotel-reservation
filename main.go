package main

import (
	"flag"
	"fmt"

	"github.com/AyanokojiKiyotaka8/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddress := flag.String("listenAddress", ":3000", "The listen address of API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	fmt.Println("Starting at port:3000")
	app.Listen(*listenAddress)
}