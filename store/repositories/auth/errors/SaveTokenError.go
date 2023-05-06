package errors

import "fmt"

type SaveTokenError struct {
	Phone         string
	Token         string
	DurationHours uint
	Reason        error
}

func (ste SaveTokenError) Error() string {
	return fmt.Sprintf("Error saving token %v for %v", ste.Phone, ste.Token)
}

func NewSaveTokenError(phone string, token string, durationHours uint, error error) error {
	return SaveTokenError{
		Phone:         phone,
		Token:         token,
		DurationHours: durationHours,
		Reason:        error,
	}
}
