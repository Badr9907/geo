package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

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
// gettting api for dates and fetching it
func fetchDates(id int) []string {
	DatesUrl := "https://groupietrackers.herokuapp.com/api/dates/"

	resp, err := http.Get(DatesUrl + strconv.Itoa(id))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var dates reeldata
	errr := json.Unmarshal(body, &dates)
	if errr != nil {
		return nil
	}

	return dates.Dates
}
