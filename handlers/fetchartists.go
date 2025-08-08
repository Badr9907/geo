package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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
type Locations struct {
	Locations []string `json:"locations"`
}
type Dates struct {
	Dates []string `json:"dates"`
}

type ArtistPageData struct {
	Artist   Artist
	Concerts []string
	Dates    []string
	Relation map[string][]string
}

type Relations struct {
	Relation map[string][]string `json:"datesLocations"`
}

func FetchArtists() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var artists []Artist
	err = json.Unmarshal(body, &artists)
	return artists, err
}

func Fetchlocation(id int) []string {
	relationsURL := "https://groupietrackers.herokuapp.com/api/locations/"
	resp, err := http.Get(relationsURL + strconv.Itoa(id))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var rel Locations
	json.Unmarshal(body, &rel)

	return rel.Locations
}

// getting api for dates and fetching it
func FetchDates(id int) []string {
	DatesUrl := "https://groupietrackers.herokuapp.com/api/dates/"

	resp, err := http.Get(DatesUrl + strconv.Itoa(id))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var dates Dates
	errr := json.Unmarshal(body, &dates)
	if errr != nil {
		return nil
	}

	return dates.Dates
}

func fetchrelation(ids int) map[string][]string {
	DatesUrl := "https://groupietrackers.herokuapp.com/api/relation/"

	resp, err := http.Get(DatesUrl + strconv.Itoa(ids))
	if err != nil {
		return map[string][]string{}
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var rel Relations
	errr := json.Unmarshal(body, &rel)
	if errr != nil {
		return map[string][]string{}
	}
	return rel.Relation
}
