package handlers

import (
	"html/template"
	"net/http"
	"strconv"
)

// HandleError renders an error page with the given error message and HTTP status code.
func HandleError(w http.ResponseWriter, errorText string, statusCode int) {
	myMap := make(map[string]string)
	myMap["errorText"] = errorText
	myMap["statusCode"] = strconv.Itoa(statusCode)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error!", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	execError := tmpl.Execute(w, myMap)
	if execError != nil {
		http.Error(w, "500 Internal Server Error!", http.StatusInternalServerError)
		return
	}
}