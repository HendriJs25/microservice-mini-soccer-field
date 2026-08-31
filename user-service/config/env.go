package config

import (
	"fmt"
	"os"
	"strconv"
)

func getEnv(key, fallback string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) (int, error) {
	value, exist := os.LookupEnv(key)

	if !exist {
		return fallback, nil
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("failed to convert %s to int", value)
	}

	return valueInt, nil
}

func validatePort(key, value string) error {
	port, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("%s must be a number: %w", key, err)
	}

	if port < 1 || port > 65535 {
		return fmt.Errorf("%s must be between 1 and 65535", key)
	}

	return nil
}
