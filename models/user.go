package models

import "time"

type User struct {
	ID        int        `gorm:"primary_key"`
	UUID      string     `json:"-"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index"json:"-"`

	Pets []Pet
}
