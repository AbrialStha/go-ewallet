package account

import (
	"github.com/abrialstha/go-ewallet/domain"
)

type AccountStatus string

const (
	Active   AccountStatus = "ACTIVE"
	Disabled AccountStatus = "DISABLED"
	InReview AccountStatus = "IN_REVIEW"
)

type Account struct {
	Id            uint64        `json:"id"`
	AccountNumebr string        `json:"account_number"` //! Some unique value based on our app needs
	Balance       float64       `json:"balance"`
	Status        AccountStatus `json:"status"`
	UserId        string        `json:"user_id"`
	// Basic Fields
	domain.Base
}

type CreateAccount struct {
	AccountNumber string        `json:"account_number"`
	Status        AccountStatus `json:"status"`
	UserId        string        `json:"user_id"`
}
