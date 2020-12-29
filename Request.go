package util

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func ReadBody(request *http.Request) string {
	requestBody := request.Body
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(requestBody)
	return buf.String()
}

func ReadBodyToStruct(request *http.Request, obj interface{}) error {
	requestBody := request.Body
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(requestBody)
	return json.Unmarshal(buf.Bytes(), obj)
}
