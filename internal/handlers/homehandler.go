package handlers

import (
	"html/template"
	"log"
	"net/http"

	"learn.zone01kisumu.ke/git/quochieng/groupie-tracker/internal/api"
)

// HomeHandler handles the requests to the home page. It checks if the URL path is "/"
// and ensures the request method is GET. If not, it returns appropriate error messages.
// It then fetches data via the api.FetchAllData() function and renders the template with
// the fetched data. If any error occurs during these steps, it logs the error and displays
// the appropriate error messages to the user.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	homepath := "/"
	if r.URL.Path != homepath {
		HandlingErrors(w, "Page Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		HandlingErrors(w, "Method Not allowed", http.StatusMethodNotAllowed)
		return
	}
	combinedData, err := api.FetchAllData()
	if err != nil {
		log.Printf("Failed to Fetch artist: %s", err)
		HandlingErrors(w, "Problem ocurred while fetching artists, try again later", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/layout.html", "web/templates/index.html")
	if err != nil {
		log.Printf("Failed to parse file web/templates/index.html: %s", err)
		HandlingErrors(w, "File Not found", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, combinedData)
	if err != nil {
		log.Printf("Failed to execute web/templates/index.html: %s", err)
		HandlingErrors(w, "Error occured; try again later", http.StatusInternalServerError)
	}
}
