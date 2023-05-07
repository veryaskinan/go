package authorize

import (
	"main/domen/modules/auth/authorize/errors"
	authRepository "main/store/repositories/auth"
)

func Authorize(authToken string) (string, error) {
	userInfo, err := authRepository.GetByToken(authToken)
	if err != nil {
		return "", errors.NewAuthorizeError(authToken, err)
	}
	return userInfo, nil
}
