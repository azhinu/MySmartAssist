package main
import (
  "github.com/joho/godotenv"
  "os"
  "log"
)

type BotConfig struct {
    Token     string
    WitToken  string
    URL       string
    Listener  string
}


// New returns a new Config struct
func SetEnv() *BotConfig {
  if err := godotenv.Load(); err != nil {
      log.Println("No .env file found.")
  }
  return &BotConfig{
    Token: getEnv("TELEGRAM_TOKEN"),
    WitToken: getEnv("WIT_TOKEN"),
    URL: getEnv("WEBHOOK_URL"),
    Listener: getEnv("LISTEN"),
  }
}

// Simple helper function to read an environment or return a default value
func getEnv(key string) (string) {
    if value, exists := os.LookupEnv(key); exists {
	return value
    }
    panic(key+"is not defined nether .env file nor user env. ")

}
