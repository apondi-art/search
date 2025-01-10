package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"learn.zone01kisumu.ke/git/quochieng/groupie-tracker/internal/api"
	"learn.zone01kisumu.ke/git/quochieng/groupie-tracker/internal/models"
)

// ArtistsHandler handles requests to the /artist/ path. It ensures the correct path and
// request method, extracts the artist ID from the query parameters, and converts it to an integer.
// It fetches the artist data from the API, searches for the artist by ID, and renders the
// artist.html template if the artist is found. If any errors occur (e.g., invalid artist ID,
// artist not found, or issues with fetching data), appropriate error messages are sent to the client
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	pathartist := "/artist/"
	if r.URL.Path != pathartist {
		HandlingErrors(w, "Path Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		HandlingErrors(w, "Method Not allowed", http.StatusMethodNotAllowed)
		return
	}

	artistID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(artistID)
	if err != nil {
		log.Printf("Error converting artist ID: %v", err)
		HandlingErrors(w, "Artist Not Found", http.StatusNotFound)
		return
	}

	artists, err := api.FetchAllData()
	if err != nil {
		log.Printf("Error fetching artist data: %v", err)
		HandlingErrors(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var selectedArtist *models.CombinedData
	for i := range artists {
		if artists[i].Artist.ID == id {
			selectedArtist = &artists[i]
			break
		}
	}

	if selectedArtist == nil {
		HandlingErrors(w, "Artist Not Found", http.StatusNotFound)
		return
	}

	artistTemplate, err := template.ParseFiles("web/templates/layout.html", "web/templates/artist.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		HandlingErrors(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = artistTemplate.Execute(w, selectedArtist)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		HandlingErrors(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
