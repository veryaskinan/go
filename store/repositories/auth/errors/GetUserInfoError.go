package errors

import "fmt"

type GetUserInfoError struct {
	Token  string
	Reason error
}

func (guie GetUserInfoError) Error() string {
	return fmt.Sprintf("Error getting user info by token %v", guie.Token)
}

func NewGetUserInfoError(token string, error error) error {
	return GetUserInfoError{
		Token:  token,
		Reason: error,
	}
}
