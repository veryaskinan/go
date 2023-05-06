package errors

import "fmt"

type GetError struct {
	Key    string
	Reason error
}

func (r GetError) Error() string {
	return fmt.Sprintf("Error getting %v", r.Key)
}

func NewGetError(key string, error error) error {
	return GetError{
		Key:    key,
		Reason: error,
	}
}
