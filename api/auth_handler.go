package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AyanokojiKiyotaka8/hotel-reservation/db"
	"github.com/AyanokojiKiyotaka8/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	userStore db.UserStore
}

func NewAuthHandler(userStore db.UserStore) *AuthHandler {
	return &AuthHandler{
		userStore:userStore,
	}
}

type AuthParams struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}

type AuthResponse struct {
	User 	*types.User `json:"user"`
	Token 	string 		`json:"token"`
}

type genericResp struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

func invalidCredentials(c *fiber.Ctx) error {
	return c.Status(http.StatusBadRequest).JSON(genericResp{
		Type: "error",
		Msg: "invalid credentials",
	})
}

func (h *AuthHandler) HandleAuthenticate(c *fiber.Ctx) error {
	var params AuthParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	user, err := h.userStore.GetUserByEmail(c.Context(), params.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return invalidCredentials(c)
		}
		return err
	}

	if !types.IsValidPassword(user.EncryptedPassword, params.Password) {
		return invalidCredentials(c)
	}
	
	resp := AuthResponse{
		User: user,
		Token: makeClaimsFromUser(user),
	}
	return c.JSON(resp)
}

func makeClaimsFromUser(user *types.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"email": user.Email,
		"expires": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println("failed to sign token with secret", err)
	}
	return tokenString
}