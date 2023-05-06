package request

import (
	authRepository "main/store/repositories/auth"
	"main/utils/generate"
)

func Request(requestInfo RequestInfo) (string, error) {
	code := authRepository.GetAuthCode(requestInfo.Phone)

	if code == "" {
		code = generateCode()
		authRepository.SaveAuthCode(requestInfo.Phone, code, 1)
	}

	return code, nil
}

func generateCode() string {
	return generate.Number(4)
}
