package request

import (
	"errors"
	authRepository "main/store/repositories/auth"
	"main/utils/generate"
)

func Request(requestInfo RequestInfo) (string, error) {
	code := authRepository.GetAuthCode(requestInfo.Phone)

	if code == "" {
		code, _ = generateCode()
	}

	authRepository.SaveAuthCode(requestInfo.Phone, code, 3)

	return code, nil
}

func generateCode() (string, error) {
	code, err := generate.Number(4)
	if err != nil {
		return "", errors.New("Signup error: error while generating code")
	}
	return code, nil
}
