package routes

import (
	"github.com/felipe-seo/br-api/src/models"
	"github.com/gorilla/mux"
)

func RouterSetup(r *mux.Router) *mux.Router {
	routes := models.Router
	return r
}
