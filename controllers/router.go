package controllers

import (
	httpServerTypes "main/_lib/httpServer/types"
	"main/controllers/auth"
	"main/controllers/welcome"
)

func GetEnpoints() []httpServerTypes.Endpoint {
	endpoints := []httpServerTypes.Endpoint{}

	endpoints = append(endpoints, welcome.GetEndpoints()...)
	endpoints = append(endpoints, auth.GetEndpoints()...)

	return endpoints
}
