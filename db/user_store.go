package db

import "github.com/AyanokojiKiyotaka8/hotel-reservation/types"

type UserStore interface {
	GetUserByID(string) (*types.User, error)
}

type MongoUserStore struct {}
