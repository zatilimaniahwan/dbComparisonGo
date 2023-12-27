package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

const DEV_ENV = "dev.env"
const STAG_ENV = "stag.env"
const PROD_ENV = "prod.env"

func HandleLoadEnv(environment string) string {
	switch environment {
	case "stag":
		return STAG_ENV
	case "prod":
		return PROD_ENV
	default:
		return DEV_ENV
	}
}

func ConnectDB(environment string) *gorm.DB {
	var envs map[string]string

	envs, err := godotenv.Read(HandleLoadEnv(environment))
	if err != nil {
		panic("Error loading .env file")
	}
	DB_USERNAME := envs["DB_USERNAME"]
	DB_PASSWORD := envs["DB_PASSWORD"]
	DB_NAME := envs["DB_NAME"]
	DB_HOST := envs["DB_HOST"]
	DB_PORT := envs["DB_PORT"]

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		println("connected to database")
	}
	Db = db
	return Db
}
