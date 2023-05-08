package welcome

import (
	httpServerTypes "main/_lib/httpServer/types"
	"main/represent/middlewares/auth"
)

func GetEndpoints() []httpServerTypes.Endpoint {
	return []httpServerTypes.Endpoint{
		{Uri: "/", Handler: handler, Middlewares: []httpServerTypes.Middleware{auth.AuthMiddleware}},
	}
}
