package models

import "time"

type Food struct {
	Model
	PetUUID   string    `json:"pet_uuid"`
	Contents  string    `json:"contents"`
	Category  string    `json:"category"`
	StartDate time.Time `json:"start_date"`
	Url       string    `json:"url"`
	Memo      string    `json:"memo"`
}
