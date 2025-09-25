package handlers

import (
	"html/template"
	"net/http"
)

type IndexPageData struct {
	Artists []Artist
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandleError(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		HandleError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	artists, err := FetchArtists()
	if err != nil {
		HandleError(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	filter := ParseFilters(r.URL.Query())
	filtered := FilterArtists(artists, filter, nil)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		HandleError(w, "can't access this file", http.StatusInternalServerError)
		return
	}
	data := IndexPageData{
		Artists: filtered,
	}
	if err := tmpl.Execute(w, data); err != nil {
		HandleError(w, "Failed to show artists", http.StatusInternalServerError)
	}
}
