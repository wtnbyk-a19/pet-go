package model

import "time"

type Salon struct {
	Model
	PetID       string    `json:"pet_id"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"start_date"`
	Status      string    `json:"status"`
	PhoneNumber string    `json:"phine_number"`
	PostalCode  string    `json:"postal_code"`
	Address     string    `json:"address"`
	Url         string    `json:"url"`
	Memo        string    `json:"memo"`
}
