package auth

import httpServerTypes "main/_lib/httpServer/types"
import "main/controllers/auth/signIn/request"
import "main/controllers/auth/signIn/confirm"

func GetEndpoints() []httpServerTypes.Endpoint {
	return []httpServerTypes.Endpoint{
		{Uri: "/auth/signIn/request", Handler: request.Handler},
		{Uri: "/auth/signIn/confirm", Handler: confirm.Handler},
	}
}
