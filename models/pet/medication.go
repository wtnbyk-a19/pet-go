package pet

import "time"

type Medication struct {
	ID        int        `gorm:"primary_key"`
	Contents  string     `json:"contents"`
	Category  string     `json:"category"`
	Date      time.Time  `json:"date"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index"json:"-"`
}
