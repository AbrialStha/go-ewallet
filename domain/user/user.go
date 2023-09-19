package user

import (
	"time"

	"github.com/abrialstha/go-ewallet/domain"
)

type User struct {
	Id          uint64 `json:"id"`
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Location    string `json:"location"` //! Later we may want this to be a struct of some type
	Age         string `json:"age"`
	Gender      string `json:"gender"`
	// Basic Fields
	domain.Base
}

type CreateUser struct {
	FirstName   string    `json:"first_name"`
	MiddleName  string    `json:"middle_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserLogin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}
