package pet

import "time"

type Food struct {
	ID        int        `gorm:"primary_key"`
	Contents  string     `json:"contents"`
	Category  string     `json:"category"`
	StartDate time.Time  `json:"start_date"`
	Url       string     `json:"url"`
	Memo      string     `json:"memo"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index"json:"-"`
}
