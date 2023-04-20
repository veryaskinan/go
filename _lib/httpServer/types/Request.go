package types

import (
	"encoding/json"
	"io"
	"net/http"
)

type AnyMap map[string]interface{}

type Request struct {
	Request *http.Request
}

func (r Request) GetBody() AnyMap {
	bodyBytes, _ := io.ReadAll(r.Request.Body)
	var bodyMap AnyMap
	_ = json.Unmarshal(bodyBytes, &bodyMap)
	return bodyMap
}
