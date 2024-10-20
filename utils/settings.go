package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type Settings struct {
	JWTSecret string
}

var appettings *Settings

func GetSettings() *Settings {
	if appettings == nil {
		appettings = InitSettings()
	}
	return appettings
}

func InitSettings() *Settings {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return &Settings{
		JWTSecret: ReadFromEnv("JWT_SECRET"),
	}
}

func ReadFromEnv(key string) string {
	value := os.Getenv(key)
	return value
}
