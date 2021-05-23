package model

import "time"

type Food struct {
	Model
	PetID     string    `json:"pet_id"`
	Contents  string    `json:"contents"`
	Category  string    `json:"category"`
	StartDate time.Time `json:"start_date"`
	Url       string    `json:"url"`
	Memo      string    `json:"memo"`
}
