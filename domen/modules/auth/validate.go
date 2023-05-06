package auth

import (
	"errors"
	"main/_lib/logger"
	authRepository "main/store/repositories/auth"
)

var authValidationError = errors.New("AuthValidationError")

func Validate(authToken string) (string, error) {
	userInfo, err := authRepository.GetByToken(authToken)
	if err != nil {
		if err.Error() == "tokenNotFound" {
			logger.Info("ERROR! auth.Validate() token not found")
		}
		return "", authValidationError
	}
	return userInfo, nil
}
