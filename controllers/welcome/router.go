package welcome

import (
	httpServerTypes "main/_lib/httpServer/types"
)

func GetEndpoints() []httpServerTypes.Endpoint {
	return []httpServerTypes.Endpoint{
		{Uri: "/", Handler: handler},
	}
}
