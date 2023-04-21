package httpServer

import (
	"main/_lib/httpServer/types"
	"main/_lib/logger"
	"net/http"
)

type LogMessage struct {
	Time    string      `json:"time"`
	Level   string      `json:"level"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func handle(uri string, handler func(req types.Request, res types.Response)) {
	http.HandleFunc(uri, func(rw http.ResponseWriter, req *http.Request) {
		httpRequest := types.Request{req}
		httpResponse := types.Response{rw}
		handler(httpRequest, httpResponse)
	})
}

func Start(port string, endpoints []types.Endpoint) {
	logger.Info("HTTP SERVER: Init endpoints")

	for _, endpoint := range endpoints {
		handle(endpoint.Uri, endpoint.Handler)
		logger.Info("HTTP SERVER: Init endpoint " + endpoint.Uri)
	}

	logger.Info("HTTP SERVER: All endpoints inited")

	logger.Info("HTTP SERVER: Listen " + port)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		logger.Info("ERROR! Http server start error")
	}
}
