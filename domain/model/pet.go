package model

import (
	"time"
)

type Pet struct {
	Model
	Name          string    `json:"name"`
	Gender        string    `json:"gender"`
	Category      string    `json:"category"`
	Breed         string    `json:"breed"`
	Adoptaversary time.Time `json:"adoptaversary"`
	Birthday      time.Time `json:"birthday"`
	Memo          string    `json:"memo"`

	DailyMeals  DailyMeals   `gorm:"foreignKey:PetUUID"`
	Foods       []Food       `gorm:"foreignKey:PetUUID"`
	Hospitals   []Hospital   `gorm:"foreignKey:PetUUID"`
	Medications []Medication `gorm:"foreignKey:PetUUID"`
	Physique    Physique     `gorm:"foreignKey:PetUUID"`
	Salons      []Salon      `gorm:"foreignKey:PetUUID"`
	Spots       []Spot       `gorm:"foreignKey:PetUUID"`
}
