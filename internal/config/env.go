package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Exit application when fail to load env file
func LoadEnv() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalln("Error while loading env file:", err.Error())
	}
}

func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
