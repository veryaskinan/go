package errors

import "fmt"

type AuthorizeError struct {
	AuthToken string
	Reason    error
}

func (ae AuthorizeError) Error() string {
	return fmt.Sprintf("Error authorize by token %v", ae.AuthToken)
}

func NewAuthorizeError(token string, error error) error {
	return AuthorizeError{
		AuthToken: token,
		Reason:    error,
	}
}
