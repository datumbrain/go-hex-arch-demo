package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	HttpPort = "8080"
	DbUrl    string
)

func Setup() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	if x := os.Getenv("HTTP_PORT"); x != "" {
		HttpPort = x
		log.Printf("HTTP_PORT is set to %s", HttpPort)
	} else {
		log.Printf("HTTP_PORT is not set, defaulting to %s", HttpPort)
	}

	if x := os.Getenv("DATABASE_URL"); x != "" {
		DbUrl = x
		log.Println("DATABASE_URL is set")
	} else {
		log.Fatal("DATABASE_URL is not set")
	}
}
