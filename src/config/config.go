package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConnectionDB = ""
	Port               = 0
	SecretKey          []byte
)

func LoadConfig() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 5000
	}

	StringConnectionDB = fmt.Sprintf("%s:%s@/%s?allowPublicKeyRetrieval=true&useSSL=false",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}