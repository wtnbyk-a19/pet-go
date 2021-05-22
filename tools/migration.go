package main

import (
	"pet-go/domain/models"
	"pet-go/infrastructure/databases"

	"github.com/sirupsen/logrus"
)

func main() {
	db, err := databases.Connect()
	defer db.Close()

	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Pet{})
	db.Debug().AutoMigrate(&models.DailyMeals{})
	db.Debug().AutoMigrate(&models.Food{})
	db.Debug().AutoMigrate(&models.Hospital{})
	db.Debug().AutoMigrate(&models.Medication{})
	db.Debug().AutoMigrate(&models.Physique{})
	db.Debug().AutoMigrate(&models.Salon{})
	db.Debug().AutoMigrate(&models.Spot{})
}
