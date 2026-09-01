package config

import (
	"fmt"
	"strings"
)

type Seed struct {
	AdminPassword string
	AdminEmail    string
}

func loadSeed() Seed {
	return Seed{
		AdminPassword: getEnv("SEED_ADMIN_PASSWORD", ""),
		AdminEmail:    getEnv("SEED_ADMIN_EMAIL", ""),
	}
}

func (s *Seed) Validate() error {
	if strings.TrimSpace(s.AdminEmail) == "" {
		return fmt.Errorf("SEED_ADMIN_EMAIL must not be empty")
	}

	if strings.TrimSpace(s.AdminPassword) == "" {
		return fmt.Errorf("SEED_ADMIN_PASSWORD must not be empty")
	}

	if len(s.AdminPassword) < 8 {
		return fmt.Errorf("SEED_ADMIN_PASSWORD must not be less than 8 characters")
	}

	if len(s.AdminPassword) > 72 {
		return fmt.Errorf("SEED_ADMIN_PASSWORD must not be more than 72 characters")
	}
	return nil
}
