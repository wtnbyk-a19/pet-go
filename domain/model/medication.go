package model

import "time"

type Medication struct {
	Model
	PetID    string    `json:"pet_id"`
	Contents string    `json:"contents"`
	Category string    `json:"category"`
	Date     time.Time `json:"date"`
}
