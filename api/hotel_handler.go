package api

import (
	"fmt"

	"github.com/AyanokojiKiyotaka8/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(hotel)
}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	
	rooms, err := h.store.Room.GetRooms(c.Context(), bson.M{"hotelID": oid})
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}

type HotelQueryParams struct {
	Rooms bool
	Rating int
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var queryParams HotelQueryParams
	if err := c.QueryParser(&queryParams); err != nil {
		return err
	}
	fmt.Println(queryParams)

	hotels, err := h.store.Hotel.GetHotels(c.Context(), bson.M{})
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}