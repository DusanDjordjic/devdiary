package main

import (
	"dev-diary/router"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	router.SetupRouter(mux)

	server := http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
