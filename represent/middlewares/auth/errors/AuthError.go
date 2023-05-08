package errors

import (
	"fmt"
)

type AuthError struct {
	Reason error
}

func (au AuthError) Error() string {
	return fmt.Sprintf("Auth error")
}

func NewAuthError(error error) error {
	return AuthError{
		Reason: error,
	}
}
