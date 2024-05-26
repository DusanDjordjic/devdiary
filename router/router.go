package router

import (
	"dev-diary/router/handlers"
	"net/http"
)

func SetupRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /", handlers.HomePageHandler)
}
