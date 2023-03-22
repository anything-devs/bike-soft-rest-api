package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func VariablesEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error al cargar .env")
	}
	return os.Getenv(key)
}
