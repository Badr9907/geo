package handlers

import "net/http"

// HandlerCss serves CSS files and handles access restrictions.
func HandleStatic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/static/" || r.URL.Path == "/static" {
		HandleError(w, " Acess Forbidden", http.StatusForbidden)
		return
	}
	if r.URL.Path != "/static/style.css" {
		HandleError(w, "Page not found", http.StatusNotFound)
		return
	}

	HandlerCss := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	HandlerCss.ServeHTTP(w, r)
}
