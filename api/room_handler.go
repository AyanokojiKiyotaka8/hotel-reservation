package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AyanokojiKiyotaka8/hotel-reservation/db"
	"github.com/AyanokojiKiyotaka8/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomHandler struct {
	store *db.Store
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		store: store,
	}
}

type BookRoomParams struct {
	NumPersons			int 				`json:"numPersons,omitempty"`
	FromDate			time.Time			`json:"fromDate,omitempty"`
	TillDate			time.Time			`json:"tillDate,omitempty"`
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	roomID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}

	var params BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	user, ok := c.Context().Value("user").(*types.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(genericResp{
			Type: "error",
			Msg: "internal server error",
		})
	}

	booking := types.Booking{
		RoomID: roomID,
		UserID: user.ID,
		NumPersons: params.NumPersons,
		FromDate: params.FromDate,
		TillDate: params.TillDate,
	}

	fmt.Println(booking)
	return nil
}
