package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/pawlobanano/chat-websockets-go/internal/handlers"
)

// routes defines the application routes.
func routes() http.Handler {
	mux := pat.New()

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	return mux
}
