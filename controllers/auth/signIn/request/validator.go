package request

import (
	"errors"
	"fmt"
	httpServerTypes "main/_lib/httpServer/types"
	"main/domen/modules/auth/signIn/request"
)

func validate(req httpServerTypes.Request) (request.RequestInfo, error) {
	body := req.GetBody()

	_, phoneIsString := body["phone"].(string)

	var err error = nil

	if !phoneIsString {
		err = errors.New("Phone must be a string")
	}

	requestInfo := request.RequestInfo{
		Phone: fmt.Sprintf("%v", body["phone"]),
	}

	return requestInfo, err
}
