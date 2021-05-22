package models

type DailyMeals struct {
	Model
	PetUUID       *string `json:"pet_uuid"`
	Contents      string  `json:"contents"`
	Amount        int     `json:"amount"`
	NumberOfMeals int     `json:"number_of_meals"`
}
