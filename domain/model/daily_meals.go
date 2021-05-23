package model

type DailyMeals struct {
	Model
	PetID         string `json:"pet_id"`
	Contents      string `json:"contents"`
	Amount        int    `json:"amount"`
	NumberOfMeals int    `json:"number_of_meals"`
}
