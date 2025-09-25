package handlers

import (
	"net/url"
	"strconv"
	"strings"
)

type Filter struct {
	CreationMin int
	CreationMax int
	AlbumMin    int
	AlbumMax    int
	Members     map[int]bool
}

func ParseFilters(values url.Values) Filter {
	filter := Filter{
		CreationMin: 0,
		CreationMax: 9999,
		AlbumMin:    0,
		AlbumMax:    9999,
		Members:     make(map[int]bool),
	}
	if v := values.Get("creation_min"); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			filter.CreationMin = i
		}
	}
	if v := values.Get("creation_max"); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			filter.CreationMax = i
		}
	}
	if v := values.Get("album_min"); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			filter.AlbumMin = i
		}
	}
	if v := values.Get("album_max"); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			filter.AlbumMax = i
		}
	}
	for _, m := range values["members"] {
		if m == "6+" {
			for i := 6; i <= 20; i++ {
				filter.Members[i] = true
			}
		} else if i, err := strconv.Atoi(m); err == nil {
			filter.Members[i] = true
		}
	}
	return filter
}

func FilterArtists(artists []Artist, filter Filter, _ map[int][]string) []Artist {
	var result []Artist
	for _, a := range artists {
		if a.CreationDate < filter.CreationMin || a.CreationDate > filter.CreationMax {
			continue
		}
		albumYear, _ := strconv.Atoi(strings.Split(a.FirstAlbum, "-")[0])
		if albumYear < filter.AlbumMin || albumYear > filter.AlbumMax {
			continue
		}
		if len(filter.Members) > 0 && !filter.Members[len(a.Members)] {
			continue
		}
		result = append(result, a)
	}
	return result
}
