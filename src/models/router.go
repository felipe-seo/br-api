package models

import "net/http"

type Router struct {
	Uri      string
	Function func(w http.ResponseWriter, r *http.Request)
	Method   string
}
