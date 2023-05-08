package request

import (
	"encoding/json"
	"io"
	"main/_lib/httpServer/request/errors"
	"net/http"
)

type AnyMap map[string]interface{}

type Request struct {
	Request *http.Request
	User    string
}

func (r *Request) GetBody() AnyMap {
	bodyBytes, _ := io.ReadAll(r.Request.Body)
	var bodyMap AnyMap
	_ = json.Unmarshal(bodyBytes, &bodyMap)
	return bodyMap
}

func (r *Request) GetHeader(headerName string) (string, error) {
	header := r.Request.Header.Get(headerName)
	if header == "" {
		return "", errors.NewGetHeaderError(headerName, nil)
	}
	return header, nil
}
