package main

import (
	"pet-go/domain/model"
	"pet-go/infrastructure/database"

	"github.com/sirupsen/logrus"
)

func main() {
	db, err := database.Connect()
	defer db.Close()

	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&model.User{})
	db.Debug().AutoMigrate(&model.Pet{})
	db.Debug().AutoMigrate(&model.DailyMeals{})
	db.Debug().AutoMigrate(&model.Food{})
	db.Debug().AutoMigrate(&model.Hospital{})
	db.Debug().AutoMigrate(&model.Medication{})
	db.Debug().AutoMigrate(&model.Physique{})
	db.Debug().AutoMigrate(&model.Salon{})
	db.Debug().AutoMigrate(&model.Spot{})
}
