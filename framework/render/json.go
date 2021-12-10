package render

import "net/http"

type Json struct {
	Data interface{}
}


var jsonContentType = []string{"application/json; charset=utf-8"}

func (r Json) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}