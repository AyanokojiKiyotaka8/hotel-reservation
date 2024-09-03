package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const(
	minFirstNameLen = 2
	minLastNameLen = 2
	minPasswordLen = 7
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (p CreateUserParams) Validate() []string {
	errors := []string{}
	if len(p.FirstName) < minFirstNameLen {
		errors = append(errors, fmt.Sprintf("firstName length should be at least %d characters", minFirstNameLen))
	}
	if len(p.LastName) < minLastNameLen {
		errors = append(errors, fmt.Sprintf("lastName length should be at least %d characters", minLastNameLen))
	}
	if len(p.Password) < minPasswordLen {
		errors = append(errors, fmt.Sprintf("password length should be at least %d characters", minPasswordLen))
	}
	if !isValidEmail(p.Email) {
		errors = append(errors, fmt.Sprintf("email is invalid"))
	}
	return errors
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}


type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName string `bson:"lastName" json:"lastName"`
	Email string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"EncryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName: params.FirstName,
		LastName: params.LastName,
		Email: params.Email,
		EncryptedPassword: string(enc),
	}, nil
}