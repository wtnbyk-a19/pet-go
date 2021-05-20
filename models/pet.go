package models

import (
	"time"
)

type Pet struct {
	ID            int        `gorm:"primary_key"`
	UUID          string     `json:"uuid"`
	UserUUID      string     `json:"user_uuid"`
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

	DailyMeals  DailyMeals   `gorm:"foreignKey:PetUUID"`
	Foods       []Food       `gorm:"foreignKey:PetUUID"`
	Hospitals   []Hospital   `gorm:"foreignKey:PetUUID"`
	Medications []Medication `gorm:"foreignKey:PetUUID"`
	Physique    Physique     `gorm:"foreignKey:PetUUID"`
	Salons      []Salon      `gorm:"foreignKey:PetUUID"`
	Spots       []Spot       `gorm:"foreignKey:PetUUID"`
}
