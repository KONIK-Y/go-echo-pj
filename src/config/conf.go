package cnf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found. Proceeding with environment variables.")
    }
}


func LoadEnv(key, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultVal
}
