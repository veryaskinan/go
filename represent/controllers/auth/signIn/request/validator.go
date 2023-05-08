package request

import (
	"errors"
	"fmt"
	httpServerTypes "main/_lib/httpServer/request"
	"main/domen/modules/auth/signIn/request/types"
)

func validate(req httpServerTypes.Request) (types.RequestInfo, error) {
	body := req.GetBody()

	_, phoneIsString := body["phone"].(string)

	var err error = nil

	if !phoneIsString {
		err = errors.New("Phone must be a string")
	}

	requestInfo := types.RequestInfo{
		Phone: fmt.Sprintf("%v", body["phone"]),
	}

	return requestInfo, err
}
