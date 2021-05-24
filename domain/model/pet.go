package model

import (
	"time"
)

type Pet struct {
	Model
	UserID        int       `json:"user_id"`
	Name          string    `json:"name"`
	Gender        string    `json:"gender"`
	Category      string    `json:"category"`
	Breed         string    `json:"breed"`
	Adoptaversary time.Time `json:"adoptaversary"`
	Birthday      time.Time `json:"birthday"`
	Memo          string    `json:"memo"`

	DailyMeals  DailyMeals
	Foods       []Food
	Hospitals   []Hospital
	Medications []Medication
	Physique    Physique
	Salons      []Salon
	Spots       []Spot
}
