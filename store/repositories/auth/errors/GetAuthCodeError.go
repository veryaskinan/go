package errors

import "fmt"

type GetAuthCodeError struct {
	Phone  string
	Reason error
}

func (gace GetAuthCodeError) Error() string {
	return fmt.Sprintf("Error getting auth code for %v", gace.Phone)
}

func NewGetAuthCodeError(phone string, error error) error {
	return GetAuthCodeError{
		Phone:  phone,
		Reason: error,
	}
}
