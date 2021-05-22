package models

import "time"

type Medication struct {
	Model
	PetUUID  string    `json:"pet_uuid"`
	Contents string    `json:"contents"`
	Category string    `json:"category"`
	Date     time.Time `json:"date"`
}
