package errors

import "fmt"

type AuthCodeNotFoundError struct {
	Phone  string
	Reason error
}

func (acnfe AuthCodeNotFoundError) Error() string {
	return fmt.Sprintf("Error getting auth code for %v. Not found.", acnfe.Phone)
}

func NewAuthCodeNotFoundError(phone string, error error) error {
	return AuthCodeNotFoundError{
		Phone:  phone,
		Reason: error,
	}
}
