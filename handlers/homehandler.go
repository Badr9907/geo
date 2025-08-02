package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := FetchArtists()
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, artists)
}

func FetchArtists() ([]Artist, error) {
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

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
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

	relationsURL := selected.Relations
	relationsURL = "https://groupietrackers.herokuapp.com/api/locations"

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
	fmt.Println(newdates)

	// fmt.Println(id)
	// log.Println(dates)
	// log.Println("Artist ID:", id, "Locations:", locations,"dates",dates)

	data := ArtistPageData{
		Artist:   selected,
		Concerts: locations,
		Dates:    newdates,
	}

	tmpl := template.Must(template.ParseFiles("templates/artist.html"))
	tmpl.Execute(w, data)
}

func fetchConcerts(url string, id int) []string {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var rel RelationResponse
	json.Unmarshal(body, &rel)

	for _, entry := range rel.Index {
		if entry.ID == id {
			//	fmt.Println("4")
			log.Println(entry.Locations)
			return entry.Locations

		}
	}
	return nil
}

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
		Locations []string `json:"locations"`
	} `json:"index"`
}
type reeldata struct {
	Dates []string `json:"dates"`
}

type ArtistPageData struct {
	Artist   Artist
	Concerts []string
	Dates    []string
}

func fetchDates(id int) []string {
	DatesUrl := "https://groupietrackers.herokuapp.com/api/dates/"
	fmt.Println("dsad")

	resp, err := http.Get(DatesUrl + strconv.Itoa(id))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var dates reeldata
	errr := json.Unmarshal(body, &dates)
	if errr != nil {
		fmt.Println("dazzzzz")
	}

	return dates.Dates
}
