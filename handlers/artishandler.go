package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// a function to handle /artist path
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		HandleError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(w, "Can't find this ID", http.StatusNotFound)
		return
	}

	artists, err := FetchArtists()
	if err != nil {
		HandleError(w, "Failed to fetch artist", http.StatusInternalServerError)
		return
	}
	if id > len(artists) || id < 1 {
		HandleError(w, "Can't find this ID", http.StatusNotFound)
		return
	}

	selected := artists[id-1]
	locations := Fetchlocation(id)
	dates := FetchDates(id)
	relations := fetchrelation(id)

	data := ArtistPageData{
		Artist:   selected,
		Concerts: locations,
		Dates:    dates,
		Relation: relations,
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		HandleError(w, "can't access this file", http.StatusInternalServerError)
		return
	}

	error1 := tmpl.Execute(w, data)
	if error1 != nil{
		HandleError(w,"Failed to show artists",http.StatusInternalServerError)
	}
}
