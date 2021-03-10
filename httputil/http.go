package httputil

import (
	"encoding/json"
	"github.com/fire-g/mark-go-util/logger"
	"net/http"
)

//返回对象,将i对象转换成json并返回给http Response
func SendBack(writer http.ResponseWriter, i interface{}) {
	jsons, err := json.Marshal(i)
	if err != nil {
		logger.Error.Println(err)
	}
	_, _ = writer.Write(jsons)
}
