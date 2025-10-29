package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	APIUSERNAME string
	APIPASSWORD string
	APIURL      string
)

func GetEnvironmentVariables() {
	if os.Getenv("VERCEL_ENV") == "production" {
		// Load environment variables from Vercel
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading environment variables from Vercel:", err)
		}
	} else {
		// Load environment variables from local .env file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	APIUSERNAME = os.Getenv("API_USERNAME")
	APIPASSWORD = os.Getenv("API_PASSWORD")
	APIURL = os.Getenv("API_URL")

	if APIUSERNAME == "" || APIPASSWORD == "" || APIURL == "" {
		log.Fatal("Missing API credentials")
	}
}
