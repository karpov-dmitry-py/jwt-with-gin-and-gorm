package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

const (
	pathToEnvFile = ".env"
)

func LoadEnvVariables() error {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		return fmt.Errorf("failed to load env variables: %v", err)
	}

	return nil
}
