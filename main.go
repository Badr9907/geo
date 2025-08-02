package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"`
}

type RelationResponse struct {
	Index []struct {
		ID        int      `json:"id"`
		Dates     []string `json:"dates"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

type ArtistPageData struct {
	Artist   Artist
	Concerts []string
	Dates    []string
}

func fetchArtists() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var artists []Artist
	err = json.Unmarshal(body, &artists)
	return artists, err
}

func fetchConcerts(url string, id int) ([]string, []string) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var rel RelationResponse
	json.Unmarshal(body, &rel)

	for _, entry := range rel.Index {
		if entry.ID == id {
			return entry.Locations, entry.Dates
		}
	}
	return nil, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := fetchArtists()
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, artists)
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artists, err := fetchArtists()
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

	locations, dates := fetchConcerts(selected.Relations, id)

	data := ArtistPageData{
		Artist:   selected,
		Concerts: locations,
		Dates:    dates,
	}

	tmpl := template.Must(template.ParseFiles("templates/artist.html"))
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artist/", artistHandler)
	log.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
