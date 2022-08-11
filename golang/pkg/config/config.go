package config

import (
	"strings"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const (
	errNoSuchFileDirectory = "no such file or directory"
)

// PopulateFromDotEnv attempts to retrieve env vars from .env files
func PopulateFromDotEnv() error {
	err := godotenv.Load(".env.local")
	if err != nil && !isNoSuchFileErr(err) {
		zap.L().Error("Failed to load .env.local", zap.Error(err))

		return err
	}
	zap.L().Info("Loaded .env.local")

	err = godotenv.Load()
	if err != nil && !isNoSuchFileErr(err) {
		zap.L().Error("Failed to load .env", zap.Error(err))

		return err
	}

	zap.L().Info("Loaded .env")

	return nil
}

func isNoSuchFileErr(err error) bool {
	return strings.Contains(err.Error(), errNoSuchFileDirectory)
}
