package types

import (
	"encoding/json"
	"fmt"
	"main/_lib/logger"
	"net/http"
)

type Response struct {
	Writer http.ResponseWriter
}

func (r *Response) SetStatus(statusCode int) {
	r.Writer.WriteHeader(statusCode)
}

func (r *Response) SendJson(data any) {
	r.Writer.Header().Set("Content-Type", "application/json")

	jsonData, jsonError := json.Marshal(data)

	if jsonError != nil {
		logger.Info("ERROR! Json encode error")
	}

	_, err := fmt.Fprintf(r.Writer, string(jsonData))
	if err != nil {
		logger.Info("ERROR! Send response error")
	}
}

func (r *Response) SendError(errorMessage string) {
	r.Writer.Header().Set("Content-Type", "application/json")

	jsonData, jsonError := json.Marshal(errorMessage)

	if jsonError != nil {
		logger.Info("ERROR! Json encode error")
	}

	_, err := fmt.Fprintf(r.Writer, string(jsonData))

	if err != nil {
		logger.Info("ERROR! Send response error")
	}
}
