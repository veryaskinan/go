package errors

import "fmt"

type SaveAuthCodeError struct {
	Phone           string
	Code            string
	DurationMinutes uint
	Reason          error
}

func (sace SaveAuthCodeError) Error() string {
	return fmt.Sprintf("Error saving auth code %v for %v", sace.Code, sace.Phone)
}

func NewSaveAuthCodeError(phone string, code string, duration uint, error error) error {
	return SaveAuthCodeError{
		Phone:           phone,
		Code:            code,
		DurationMinutes: duration,
		Reason:          error,
	}
}
