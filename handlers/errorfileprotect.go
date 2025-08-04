package handlers

import (
	"crypto/sha256"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

const originalErrorHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <style>.errorSection {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    background: #f8f9fa;
    text-align: center;
    font-family: Arial, sans-serif;
}

.statusCode {
    font-size: 5rem;
    color: #dc3545;
    margin: 0;
}

.errorText {
    font-size: 1.2rem;
    color: #333;
    margin: 10px 0 20px;
}

.btnBack {
    text-decoration: none !important;
    padding: 10px 20px;
    background-color: #1e92ff;
    border: none;
    border-radius: 5px;
    color: white;
    font-size: 1rem;
    cursor: pointer;
    display: block;
    margin-top: 20px;
}

.btnBack:hover {
    background-color: #0056b3;
}</style>
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


const expectedHash = "b52505d7a222437f234372b993a6aba828a836ad9a2df14557aaabbbecdd5015" 

func getFileHash(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash), nil
}


func restoreErrorHTML(path string) error {
	return os.WriteFile(path, []byte(originalErrorHTML), 0644)
}


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
