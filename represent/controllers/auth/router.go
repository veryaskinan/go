package auth

import (
	httpServerTypes "main/_lib/httpServer/types"
	"main/represent/controllers/auth/signIn/confirm"
	"main/represent/controllers/auth/signIn/request"
)

func GetEndpoints() []httpServerTypes.Endpoint {
	return []httpServerTypes.Endpoint{
		{Uri: "/auth/signIn/request", Handler: request.Handler},
		{Uri: "/auth/signIn/confirm", Handler: confirm.Handler},
	}
}
