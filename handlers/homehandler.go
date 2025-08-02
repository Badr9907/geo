package handlers

import (
	"net/http"
)

type Artist struct {
	Name         string   `json:"name"`
	id           string   `json:"id"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate string   `json:"creationdate"`
	Locations    string   `json:"locations"`
	Dates        string   `json:"dates"`
	FirstAlbum   string   `json:"firstalbum"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		http.Error(w, "Unable to fetch artists", http.StatusInternalServerError)
		return
	}
}
