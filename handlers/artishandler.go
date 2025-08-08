package handlers

import (
	//"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		HandleError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id >52 || id < 1 {
		HandleError(w, "Can't find this ID", http.StatusNotFound)
		return
	}

	artists, err := FetchArtists()
	if err != nil {
		HandleError(w, "Failed to fetch artist", http.StatusInternalServerError)
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

	locations := FetchConcerts(relationsURL, id)
	dates := FetchDates(id)
	//var newdates []string
	relations ,_ := fetchrelation(id)
	//fmt.Println(relations)


	

	data := ArtistPageData{
		Artist:   selected,
		Concerts: locations,
		Dates:    dates,
		Relation:  relations,
	}

	tmpl ,err:= template.ParseFiles("templates/artist.html")
	if err!=nil{
		HandleError(w,"can't access this file",http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}


