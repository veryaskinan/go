package errors

import "fmt"

type NotFoundError struct {
	Key    string
	Reason error
}

func (nfe NotFoundError) Error() string {
	return fmt.Sprintf("Error getting %v. Not found", nfe.Key)
}

func NewNotFoundError(key string, error error) error {
	return NotFoundError{
		Key:    key,
		Reason: error,
	}
}
