package types

import (
	"encoding/json"
	"io"
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

func (r *Request) GetHeader(headerName string) string {
	return r.Request.Header.Get(headerName)
}
