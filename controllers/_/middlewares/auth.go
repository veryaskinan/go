package middlewares

import (
	"errors"
	httpServerTypes "main/_lib/httpServer/types"
	"main/_lib/logger"
	"main/domen/modules/auth"
)

func AuthMiddleware(request httpServerTypes.Request, response httpServerTypes.Response) (httpServerTypes.Request, httpServerTypes.Response, error) {
	authToken := request.GetHeader("AUTH_TOKEN")
	userInfo, err := auth.Validate(authToken)
	if err != nil {
		if err.Error() == "AuthValidationError" {
			logger.Info("ERROR! (AuthMiddleware) Auth validation error")
			response.SendError("Auth validation error")
		}
		return request, response, errors.New("AuthValidationError")
	} else {
		request.User = userInfo
	}
	return request, response, nil
}
