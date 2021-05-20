package models

import "time"

type Salon struct {
	ID          int        `gorm:"primary_key"`
	PetUUID     string     `json:"pet_uuid"`
	Name        string     `json:"name"`
	StartDate   time.Time  `json:"start_date"`
	Status      string     `json:"status"`
	PhoneNumber string     `json:"phine_number"`
	PostalCode  string     `json:"postal_code"`
	Address     string     `json:"address"`
	Url         string     `json:"url"`
	Memo        string     `json:"memo"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `sql:"index"json:"-"`

	Pet Pet
}
