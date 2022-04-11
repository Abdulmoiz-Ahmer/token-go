package token

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidSecretKeyLen ...
	ErrInvalidSecretKeyLen = fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)

	// ErrInvalidToken ...
	ErrInvalidToken = errors.New("token is invalid")
	// ErrExpiredToken ...
	ErrExpiredToken = errors.New("token has expired")
)
