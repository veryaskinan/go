package confirm

import (
	"errors"
	"fmt"
	httpServerTypes "main/_lib/httpServer/request"
	"main/domen/modules/auth/signIn/confirm"
)

func validate(req httpServerTypes.Request) (confirm.ConfirmInfo, error) {
	body := req.GetBody()

	_, phoneIsString := body["phone"].(string)
	_, codeIsString := body["code"].(string)

	var err error = nil

	if !phoneIsString {
		err = errors.New("Phone must be a string")
	}

	if !codeIsString {
		err = errors.New("Code must be a string")
	}

	confirmInfo := confirm.ConfirmInfo{
		Phone: fmt.Sprintf("%v", body["phone"]),
		Code:  fmt.Sprintf("%v", body["code"]),
	}

	return confirmInfo, err
}
