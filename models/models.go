package main

import (
	"fmt"
	"time"
)

type user struct {
	Name      *string   `json:"Name" validate:"required,min=2,max=30"`
	Email     *string   `json:"Email" validate:"required,min=2,max=30"`
	Password  *string   `json:"Password" validate:"required, min=8, max=16"`
	Phone     *string   `json:"Phone" validate:"required"`
	CreatedAT time.Time `json:"Created At"`
	userID    string    `json:"userID"`
}

func main() {
	fmt.Println("Models will be stored here")
}
