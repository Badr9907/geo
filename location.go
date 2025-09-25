package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Coordinates struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func GeocodeLocation(address string) (Coordinates, error) {
	var coords Coordinates
	baseURL := "https://nominatim.openstreetmap.org/search"
	params := url.Values{}
	params.Set("q", address)
	params.Set("format", "json")
	params.Set("limit", "1")
	reqURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return coords, err
	}
	req.Header.Set("User-Agent", "groupie-tracker")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return coords, err
	}
	defer resp.Body.Close()

	var result []Coordinates
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return coords, err
	}
	if len(result) == 0 {
		return coords, fmt.Errorf("no coordinates found")
	}
	return result[0], nil
}
