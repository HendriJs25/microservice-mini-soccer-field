package error

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

var ErrValidator = map[string]string{
	"required": "%s is required",
	"email":    "%s is not a valid email",
	"eqfield":  "%s is not equal to %s",
}

func ErrValidationResponse(err error) (validationResponse []ValidationResponse) {
	var fieldErrors validator.ValidationErrors

	if errors.As(err, &fieldErrors) {
		for _, err := range fieldErrors {
			errValidator, ok := ErrValidator[err.Tag()]
			if ok {
				count := strings.Count(errValidator, "%s")
				if count == 1 {
					validationResponse = append(validationResponse, ValidationResponse{
						Field:   err.Field(),
						Message: fmt.Sprintf(errValidator, err.Field()),
					})
				} else {
					validationResponse = append(validationResponse, ValidationResponse{
						Field:   err.Field(),
						Message: fmt.Sprintf(errValidator, err.Field(), err.Param()),
					})
				}
			} else {
				validationResponse = append(validationResponse, ValidationResponse{
					Field:   err.Field(),
					Message: fmt.Sprintf("something wrong on %s : %s", err.Field(), err.Tag()),
				})
			}
		}
	}
	return validationResponse
}
