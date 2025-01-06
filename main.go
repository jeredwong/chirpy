package main

import (
	"net/http"
	"log"
)

func main() {
	// constants 
	const port = "8080"

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        mux,
	}

	mux.Handle("/", http.FileServer(http.Dir("")))

	log.Println("listening on port 8080...")
	log.Fatal(server.ListenAndServe())
}
