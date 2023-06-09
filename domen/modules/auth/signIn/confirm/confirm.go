package confirm

import (
	"errors"
	"main/_lib/generate"
	authRepository "main/store/repositories/auth"
	usersRepository "main/store/repositories/users"
)

func Confirm(confirmInfo ConfirmInfo) (string, error) {
	authCode, err := authRepository.GetAuthCode(confirmInfo.Phone)
	if err != nil {
		return "errors", errors.New("errors!!!!!!")
	}
	if confirmInfo.Code != authCode {
		return "errors", errors.New("Wrong auth code")
	}
	token := generate.String(16)
	go authRepository.SaveToken(confirmInfo.Phone, token, 1)
	go usersRepository.Create(confirmInfo.Phone)
	return token, nil
}
