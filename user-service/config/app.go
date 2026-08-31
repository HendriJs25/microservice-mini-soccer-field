package config

import (
	"fmt"
	"strings"
)

const (
	defaultAppEnv  = "development"
	defaultAppName = "user-service"
	defaultAppPort = "8001"
)

type App struct {
	Env          string
	Name         string
	Port         string
	SignatureKey string
}

func loadApp() App {
	return App{
		Env:          getEnv("APP_ENV", defaultAppEnv),
		Name:         getEnv("APP_NAME", defaultAppName),
		Port:         getEnv("APP_PORT", defaultAppPort),
		SignatureKey: getEnv("APP_SIGNATURE_KEY", ""),
	}
}

func (a *App) Validate() error {
	if strings.TrimSpace(a.Env) == "" {
		return fmt.Errorf("APP_ENV is required")
	}
	if strings.TrimSpace(a.Name) == "" {
		return fmt.Errorf("APP_NAME is required")
	}
	if err := validatePort("APP_PORT", a.Port); err != nil {
		return err
	}
	if strings.TrimSpace(a.SignatureKey) == "" {
		return fmt.Errorf("APP_SIGNATURE_KEY is required")
	}
	return nil
}
