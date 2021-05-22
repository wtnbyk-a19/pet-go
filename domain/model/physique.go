package model

import "time"

type Physique struct {
	Model
	PetID             string    `json:"pet_id"`
	Length            int       `json:"length"`
	Height            int       `json:"height"`
	Weight            int       `json:"weight"`
	Neckline          int       `json:"neckline"`
	Chest             int       `json:"chest"`
	Waistline         int       `json:"waistline"`
	BackLength        int       `json:"back_length"`
	LegLength         int       `json:"leg_length"`
	PawWidth          int       `json:"paw_width"`
	PawLength         int       `json:"paw_length"`
	DateOfMeasurement time.Time `json:"date_of_measurement"`
	Memo              string    `json:"memo"`
}
