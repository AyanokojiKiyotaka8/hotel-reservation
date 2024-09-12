package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AyanokojiKiyotaka8/hotel-reservation/db"
	"github.com/AyanokojiKiyotaka8/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	hotelStore *db.MongoHotelStore
	roomStore *db.MongoRoomStore
	ctx = context.Background()
)

func seedHotel(name string, location string, rating int) {
	hotel := types.Hotel{
		Name: name,
		Location: location,
		Rooms: []primitive.ObjectID{},
		Rating: rating,
	}

	rooms := []types.Room{
		{
			Size: "small",
			Price: 99.9,
			Seaside: true,
		},
		{
			Size: "kingsize",
			Price: 199.9,
			Seaside: true,
		},
		{
			Size: "normal",
			Price: 122.9,
			Seaside: false,
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

func main() {
	seedHotel("Aaaa", "aaaaaaa", 7)
	seedHotel("Bbbb", "bbbbbbb", 30)
}

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client, db.DBNAME)
	roomStore = db.NewMongoRoomStore(client, db.DBNAME, hotelStore)
}