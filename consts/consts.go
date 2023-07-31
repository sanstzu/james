package consts

import (
	"log"
	"os"

	gotdotenv "github.com/joho/godotenv"
)

func init() {
	err := gotdotenv.Load()
	log.Printf("Loaded .env file")
	if err != nil {
		log.Fatal(err)
	}
}

func ENV(key string) string {
	return os.Getenv(key)
}
