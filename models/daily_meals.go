package models

import (
	"time"
)

type DailyMeals struct {
	ID            int        `gorm:"primary_key"`
	PetUUID       string     `json:"pet_uuid"`
	Contents      string     `json:"contents"`
	Amount        int        `json:"amount"`
	NumberOfMeals int        `json:"number_of_meals"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     time.Time  `json:"-"`
	DeletedAt     *time.Time `sql:"index"json:"-"`

	Pet Pet
}
