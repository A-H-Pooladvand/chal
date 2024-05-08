package app

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"strings"
)

func Production() bool {
	env := strings.ToLower(os.Getenv("APP_ENV"))

	return env == "prod" || env == "production"
}

func Local() bool {
	return !Production()
}

func LoadEnvironmentVariablesInLocalEnv() {
	if Local() {
		if err := godotenv.Load(); err != nil {
			zap.L().Panic("unable to load .env file", zap.Error(err))
		}
	}
}
