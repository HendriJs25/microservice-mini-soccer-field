package validator

import (
	"log/slog"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
)

func New() *validator.Validate {
	v := validator.New()
	if err := v.RegisterValidation("notblank", notBlank); err != nil {
		slog.Error("Custom validator failed to register", "error", err)
		os.Exit(1)
	}
	return v
}

func notBlank(f1 validator.FieldLevel) bool {
	return strings.TrimSpace(f1.Field().String()) != ""
}
