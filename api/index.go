// Package handler is the Vercel Go serverless entrypoint.
// All HTTP traffic is rewritten to this function (see vercel.json).
package handler

import (
	"net/http"

	"nirapod/server"
)

var app = server.NewHandler()

// Handler is the entry point invoked by Vercel.
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
