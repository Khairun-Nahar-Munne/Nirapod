package handler

import (
	"net/http"

	"nirapod/server"
)

var app = server.NewHandler()

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
