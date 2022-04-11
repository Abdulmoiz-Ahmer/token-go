package token

import (
	"errors"
)

var (
	// ErrExpiredToken ...
	ErrExpiredToken = errors.New("token has expired")
)
