package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	UrlDB = ""
	Port  = 0
)

func GetVariables() {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 8080
	}

	UrlDB = os.Getenv("DATABASE_URL")
}
