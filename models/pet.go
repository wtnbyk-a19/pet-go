package models

import (
	"pet-go/models/pet"
	"time"
)

type Pet struct {
	ID            int        `gorm:"primary_key"`
	UserID        int        `json:"user_id"`
	PetName       string     `json:"pet_name"`
	Gender        string     `json:"gender"`
	Category      string     `json:"category"`
	Breed         string     `json:"breed"`
	Adoptaversary time.Time  `json:"adoptaversary"`
	Birthday      time.Time  `json:"birthday"`
	Memo          string     `json:"memo"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     time.Time  `json:"-"`
	DeletedAt     *time.Time `sql:"index"json:"-"`

	User User

	DailyMeals  pet.DailyMeals
	Foods       []pet.Food
	Hospitals   []pet.Hospital
	Medications []pet.Medication
	Physique    pet.Physique
	Salons      []pet.Salon
	Spots       []pet.Spot
}
