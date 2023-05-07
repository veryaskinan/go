package represent

import (
	httpServerTypes "main/_lib/httpServer/types"
	"main/represent/controllers/auth"
	"main/represent/controllers/welcome"
)

func GetEnpoints() []httpServerTypes.Endpoint {
	endpoints := []httpServerTypes.Endpoint{}

	endpoints = append(endpoints, welcome.GetEndpoints()...)
	endpoints = append(endpoints, auth.GetEndpoints()...)

	return endpoints
}
