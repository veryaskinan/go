package auth

import (
	"main/_lib/httpServer/request"
	httpServerTypes "main/_lib/httpServer/types"
	"main/_lib/logger"
	"main/domen/modules/auth/authorize"
	"main/represent/middlewares/auth/errors"
)

func AuthMiddleware(request request.Request, response httpServerTypes.Response) (request.Request, httpServerTypes.Response, error) {
	authToken, err := request.GetHeader("AUTH_TOKEN")

	if err != nil {
		return request, response, errors.NewAuthError(err)
	}

	userInfo, err := authorize.Authorize(authToken)

	if err != nil {
		logger.Info(err.Error())
		response.SendError(err.Error())
		return request, response, errors.NewAuthError(err)
	}

	request.User = userInfo
	return request, response, nil
}
