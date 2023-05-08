package errors

import "fmt"

type GetHeaderError struct {
	HeaderName string
	Reason     error
}

func (ghe GetHeaderError) Error() string {
	return fmt.Sprintf("Error getting header '%v'", ghe.HeaderName)
}

func NewGetHeaderError(headerName string, error error) error {
	return GetHeaderError{
		HeaderName: headerName,
		Reason:     error,
	}
}
