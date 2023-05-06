package errors

import "fmt"

type SetError struct {
	Key       string
	Value     string
	ExpireSec uint
	Reason    error
}

func (r SetError) Error() string {
	return fmt.Sprintf("Error setting %v to %v", r.Value, r.Key)
}

func NewSetError(key string, value string, expireSec uint, error error) error {
	return SetError{
		Key:       key,
		Value:     value,
		ExpireSec: expireSec,
		Reason:    error,
	}
}
