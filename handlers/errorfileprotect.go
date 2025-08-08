package handlers

import (
	"crypto/sha256"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

// The original HTML template for the error page.
// Used to restore the error.html file if it is missing or modified.
const originalErrorHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <link rel="stylesheet" href="../static/styleerror.css">
    </style>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Error!</title>
    <link rel="stylesheet" href="/static/style.css">
</head>
<body>
    <section class="errorSection">
        <h2 class="statusCode">{{.statusCode}}</h2>
        <h6 class="errorText">{{.errorText}}</h6>
        <a href="/" class="btnBack">Go to Homepage</a>
    </section>
</body>
</html>`

const expectedHash = "cf13c96145c2ee1d83dead5c76a5ebb69f1d7faab7944b81fc065b6dc581a597"

// getFileHash returns the SHA256 hash of the file at the given path.
func getFileHash(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash), nil
}

func restoreErrorHTML(path string) error {
	return os.WriteFile(path, []byte(originalErrorHTML), 0o644)
}

// HandleError checks if the error.html file exists and is unmodified.
func HandleError(w http.ResponseWriter, errorText string, statusCode int) {
	const filePath = "templates/error.html"

	currentHash, err := getFileHash(filePath)
	if os.IsNotExist(err) || currentHash != expectedHash {
		restoreErr := restoreErrorHTML(filePath)
		if restoreErr != nil {
			http.Error(w, "500 Internal Server Error (restore failed)", http.StatusInternalServerError)
			return
		}
	}
	myMap := map[string]string{
		"errorText":  errorText,
		"statusCode": strconv.Itoa(statusCode),
	}
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, "500 Internal Server Error (parse error)", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	if execErr := tmpl.Execute(w, myMap); execErr != nil {
		http.Error(w, "500 Internal Server Error (exec error)", http.StatusInternalServerError)
	}
}
