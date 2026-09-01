package error

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrTooManyRequests     = errors.New("too many requests")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrInvalidArgument     = errors.New("invalid argument")
	ErrBadRequest          = errors.New("bad request")
	ErrAlreadyExists       = errors.New("resource already exists")
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
	ErrAlreadyExists,
	ErrAccountNotVerified,
	ErrInvalidToken,
	ErrTokenExpired,
}
