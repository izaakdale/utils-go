package response

import (
	"encoding/json"
	"net/http"
)

func writeJson(writer http.ResponseWriter, code int, data interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		panic(err)
	}
}
func writeXml(writer http.ResponseWriter, code int, data interface{}) {
	writer.Header().Add("Content-Type", "application/xml")
	writer.WriteHeader(code)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		panic(err)
	}
}
