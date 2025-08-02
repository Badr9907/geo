package handlers

import (
	"html/template"
	"net/http"
)

// handling homepage request
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// checking the method
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}
	// fetching artistes then parsing files
	artists, err := FetchArtists()
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, artists)
}
