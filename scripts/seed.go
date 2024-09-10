package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AyanokojiKiyotaka8/hotel-reservation/db"
	"github.com/AyanokojiKiyotaka8/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME)
	 
	hotel := types.Hotel{
		Name: "Aaa",
		Location: "Bbb",
	}

	rooms := []types.Room{
		{
			Type: types.SingleRoomType,
			BasePrice: 99.9,
		},
		{
			Type: types.DeluxeRoomType,
			BasePrice: 199.9,
		},
		{
			Type: types.SeaSideRoomType,
			BasePrice: 122.9,
		},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(insertedHotel)

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(insertedRoom)
	}
}