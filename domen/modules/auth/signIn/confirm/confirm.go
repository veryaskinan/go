package confirm

import (
	"errors"
	authRepository "main/store/repositories/auth"
	usersRepository "main/store/repositories/users"
	"main/utils/generate"
)

func Confirm(confirmInfo ConfirmInfo) (string, error) {
	authCode := authRepository.GetAuthCode(confirmInfo.Phone)
	if confirmInfo.Code != authCode {
		return "error", errors.New("Wrong auth code")
	}
	token := generate.String(16)
	go authRepository.SaveToken(confirmInfo.Phone, token, 1)
	go usersRepository.Create(confirmInfo.Phone)
	return token, nil
}
