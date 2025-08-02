package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artists, err := FetchArtists()
	if err != nil {
		http.Error(w, "Failed to fetch artist", http.StatusInternalServerError)
		return
	}

	var selected Artist
	for _, a := range artists {
		if a.ID == id {
			selected = a
			break
		}
	}

	relationsURL := "https://groupietrackers.herokuapp.com/api/locations"

	locations := fetchConcerts(relationsURL, id)
	dates := fetchDates(id)
	var newdates []string

	for i := 0; i < len(dates); i++ {
		l := ""
		ok := dates[i]
		for j := 0; j < len(ok); j++ {
			if ok[j] != '*' {
				l += string(ok[j])
			}
		}
		newdates = append(newdates, l)

	}

	data := ArtistPageData{
		Artist:   selected,
		Concerts: locations,
		Dates:    newdates,
	}

	tmpl := template.Must(template.ParseFiles("templates/artist.html"))
	tmpl.Execute(w, data)
}
