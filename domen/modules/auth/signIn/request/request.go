package request

import (
	"main/_lib/generate"
	"main/domen/modules/auth/signIn/request/errors"
	"main/domen/modules/auth/signIn/request/types"
	authRepository "main/store/repositories/auth"
	authRepErrors "main/store/repositories/auth/errors"
)

func Request(requestInfo types.RequestInfo) (string, error) {
	code, getAuthCodeError := authRepository.GetAuthCode(requestInfo.Phone)

	if getAuthCodeError != nil {
		switch getAuthCodeError.(type) {
		case authRepErrors.AuthCodeNotFoundError:
			code = generateCode()
			saveAuthCodeError := authRepository.SaveAuthCode(requestInfo.Phone, code, 1)
			if saveAuthCodeError != nil {
				return "", errors.NewRequestError(requestInfo, saveAuthCodeError)
			}
		default:
			return "", errors.NewRequestError(requestInfo, getAuthCodeError)
		}
	}

	return code, nil
}

func generateCode() string {
	return generate.Number(4)
}
