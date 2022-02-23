package routes

import (
	"github.com/felipe-seo/br-api/src/routes/routers"
	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()
	return routers.RouterSetup(r)
}
