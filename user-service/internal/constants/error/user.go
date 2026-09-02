package error

import "errors"

var (
	ErrPasswordNotMatch       = errors.New("password does not match")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
)

var UserErrors = []error{
	ErrPasswordNotMatch,
	ErrInvalidEmailOrPassword,
}
