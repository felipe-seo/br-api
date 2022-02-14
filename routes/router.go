package routes

import "github.com/gorilla/mux"

func router() *mux.Router {
	r := mux.NewRouter()
	return r.routeSetup()
}

func routeSetup(r *mux.Router) *mux.Router {
	return r
}
