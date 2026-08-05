package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"nirapod/server"
)

func main() {
	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}

	srv := &http.Server{
		Addr:              addr,
		Handler:           server.NewHandler(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("নিরাপদ (NIRAPOD) running at http://localhost%s", addr)
	log.Fatal(srv.ListenAndServe())
}
