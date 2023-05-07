package request

import (
	"errors"
	"main/_lib/generate"
	authRepository "main/store/repositories/auth"
)

func Request(requestInfo RequestInfo) (string, error) {
	code, err := authRepository.GetAuthCode(requestInfo.Phone)

	if err != nil {
		return "", errors.New("errors!!!!!")
	}

	if code == "" {
		code = generateCode()
		authRepository.SaveAuthCode(requestInfo.Phone, code, 1)
	}

	return code, nil
}

func generateCode() string {
	return generate.Number(4)
}
