package routes

import (
	"net/http"

	"github.com/felipe-seo/br-api/src/models"
)

var endpoints = []models.Router{{
	Uri:      "/accounts",
	Function: "",
	Method:   http.MethodGet,
}}
