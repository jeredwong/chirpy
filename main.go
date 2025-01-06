package main

import (
	"log"
	"net/http"
)

func main() {
	// constants 
	const port = "8080"

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        mux,
	}

	// root path
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(""))))

	// healthz path 
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("listening on port 8080...")
	log.Fatal(server.ListenAndServe())
}
