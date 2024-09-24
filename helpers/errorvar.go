package helpers

import "errors"

var (
	ErrHashPasswordFailed     = errors.New("failed to hash password")
	ErrUserStoreFailed        = errors.New("failed to store user")
	ErrUserCheckByEmailFailed = errors.New("error checking existing user by email")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
	ErrInternalServerError    = errors.New("internal server error")
	ErrDataNotFound           = errors.New("data not found")
	ErrDuplicateEmail         = errors.New("email already registered")
)
