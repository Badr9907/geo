package handlers

import (
	"html/template"
	"net/http"
)

// handling homepage request
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"{
		HandleError(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		HandleError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// fetching artistes then parsing files
	artists, err := FetchArtists()
	if err != nil {
		HandleError(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	tmpl ,err:= template.ParseFiles("templates/index.html")
	if err != nil{
		HandleError(w,"can't access this file ",http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, artists)
}
