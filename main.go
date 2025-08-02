package main

import (
	"log"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artist/", handlers.ArtistHandler)
	log.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
