package model

import (
	"encoding/json"
	"io"
)

type Registration struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	VerifyPassword string `json:"verify_password"`
}

func RegistrationFromJson(data io.Reader) *Registration {
	var registration *Registration
	json.NewDecoder(data).Decode(&registration)
	return registration
}
