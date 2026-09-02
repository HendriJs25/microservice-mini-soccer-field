package error

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrTooManyRequests     = errors.New("too many requests")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrInvalidArgument     = errors.New("invalid argument")
	ErrBadRequest          = errors.New("bad request")
	ErrNotFound            = errors.New("not found")
	ErrAlreadyExists       = errors.New("user already exists")
	ErrAccountNotVerified  = errors.New("account not verified")

	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrForbidden,
	ErrInvalidArgument,
	ErrBadRequest,
	ErrNotFound,
	ErrAlreadyExists,
	ErrAccountNotVerified,
	ErrInvalidToken,
	ErrTokenExpired,
}
