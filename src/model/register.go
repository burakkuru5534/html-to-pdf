package model

import (
	"gopkg.in/guregu/null.v4"
)

type RegisterRequest struct {
	ID         int64       `json:"id"`
	FirstName  string      `json:"firstName" validate:"required,alphaunicode"`
	MiddleName null.String `json:"middleName" validate:"alphaunicode"`
	LastName   string      `json:"lastName" validate:"required"`
	Email      string      `json:"email" validate:"required,email"`
	Password   string      `json:"password" validate:"required,min=8"`
}

func RegisterRequestData() *RegisterRequest {
	m := new(RegisterRequest)

	return m
}
