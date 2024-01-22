package config

import(
	"log"
	"os"
	
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getEnv(key string) string {
	loadEnv()
	return os.Getenv(key)
}