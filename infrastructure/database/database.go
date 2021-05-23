package database

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() *SqlHandler {
	session, err := Connect()

	if err != nil {
		logrus.Fatal(err)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = session
	return sqlHandler
}

func Connect() (conn *gorm.DB, err error) {

	err = godotenv.Load()

	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	conn, err = gorm.Open("mysql",
		os.Getenv("DB_USERNAME")+":"+
			os.Getenv("DB_PASSWORD")+
			"@tcp("+os.Getenv("DB_HOST")+":"+
			os.Getenv("DB_PORT")+")/"+
			os.Getenv("DB_DATABASE")+
			"?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		logrus.Fatal(err)
	}

	return conn, err
}
