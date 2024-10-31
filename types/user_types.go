package types

import "time"

type CreateUserPayload struct {
	Email       string `json:"email" validate:"required,email"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	Password    string `json:"password" validate:"required,min=8"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	StreetName string `json:"streetName,omitempty"`
	StreetNumber     string `json:"streetNumber,omitempty"`
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
}

type User struct {
	ID            int       `json:"id"`
	FirstName          string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Email         string    `json:"email"`
	Password      string    `json:"-"`
	PhoneNumber   string    `json:"phone_number"`
	StreetNumber       string    `json:"streetNumber"`
	City       string    `json:"city"`
	State       string    `json:"state"`
	Country       string    `json:"Country"`
	LoyaltyPoints int       `json:"loyalty_points"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
