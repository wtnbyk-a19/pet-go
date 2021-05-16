package models

import "time"

type Pet struct {
	ID        int        `gorm:"primary_key"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index"json:"-"`

	User User
}
