package pkg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
		log.Println("Loaded .env")
		return
	}

	if _, err := os.Stat(".env.example"); err == nil {
		if err := godotenv.Load(".env.example"); err != nil {
			log.Fatalf("Error loading .env.example file: %v", err)
		}
		log.Println("Loaded .env.example")
		return
	}

	log.Fatal("No .env or .env.example file found")
}
