package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var baseUrl string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	baseUrl = os.Getenv("BASE_URL")
}
