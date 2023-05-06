package welcome

import (
	httpServerTypes "main/_lib/httpServer/types"
	middlewares "main/controllers/_/middlewares"
)

func GetEndpoints() []httpServerTypes.Endpoint {
	return []httpServerTypes.Endpoint{
		{Uri: "/", Handler: handler, Middlewares: []httpServerTypes.Middleware{middlewares.AuthMiddleware}},
	}
}
