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

	// Geocode each concert location
	var markers []MapMarker
	for _, loc := range locations {
		coords, err := GeocodeLocation(loc)
		if err == nil {
			markers = append(markers, MapMarker{
				Location: loc,
				Lat:      coords.Lat,
				Lon:      coords.Lon,
			})
		}
	}

	data := ArtistPageData{
		Artist:   selected,
		Concerts: locations,
		Dates:    dates,
		Relation: relations,
		Markers:  markers,
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		HandleError(w, "can't access this file", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		HandleError(w, "Failed to show artists", http.StatusInternalServerError)
	}
}
