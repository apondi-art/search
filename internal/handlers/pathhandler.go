package handlers

import (
	"net/http"
	"path/filepath"
)

func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		HomeHandler(w, r)
	case "/static/":
		HandlingErrors(w, "Access Forbiden", http.StatusForbidden)
	default:
		HandlingErrors(w, "Page Not Found", http.StatusNotFound)
	}
}

func CustomHandler(dir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is for the base directory
		if r.URL.Path == "/static/" {
			// Call the HandlerPath to handle the error
			PathHandler(w, r)
			return
		}

		// Remove the "/static/" prefix to get the actual file path
		filePath := filepath.Join(dir, r.URL.Path[len("/static/"):])

		http.ServeFile(w, r, filePath)
	})
}
