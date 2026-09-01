package error

import "errors"

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrUsernameAlreadyExists  = errors.New("user with this username already exists")
	ErrEmailAlreadyExists     = errors.New("user with this email already exists")
	ErrPasswordNotMatch       = errors.New("password does not match")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrUsernameAlreadyExists,
	ErrEmailAlreadyExists,
	ErrPasswordNotMatch,
	ErrInvalidEmailOrPassword,
}
