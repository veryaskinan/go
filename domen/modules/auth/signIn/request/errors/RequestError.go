package errors

import (
	"fmt"
	"main/domen/modules/auth/signIn/request/types"
)

type RequestError struct {
	RequestInfo types.RequestInfo
	Reason      error
}

func (re RequestError) Error() string {
	return fmt.Sprintf("Error signin with %v", re.RequestInfo.Phone)
}

func NewRequestError(requestInfo types.RequestInfo, error error) error {
	return RequestError{
		RequestInfo: requestInfo,
		Reason:      error,
	}
}
