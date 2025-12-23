package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file: %v", err)
		return
	}
	log.Println("env file loaded")
}

type DataBaseConfig struct {
	Url string
}

type LogConfig struct {
	Level  int
	Format string
}

func NewDatabaseConfig() *DataBaseConfig {
	return &DataBaseConfig{
		Url: getString("DATABASE_URL", "localhost"),
	}
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level:  getInt("LOG_LEVEL", 0),
		Format: getString("LOG_FORMAT", ""),
	}
}

func getString(key, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}

func getInt(key string, def int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return i
}

func getBool(key string, def bool) bool {
	value := os.Getenv(key)
	b, err := strconv.ParseBool(value)
	if err != nil {
		return def
	}
	return b
}
