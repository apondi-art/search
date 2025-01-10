package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"learn.zone01kisumu.ke/git/quochieng/groupie-tracker/internal/models"
)
// FetchDates fetches the date information for a specific artist by their ID.
// It sends a GET request to the appropriate API endpoint, decodes the JSON response,
// and returns the corresponding models.Date object or an error if something goes wrong.

var baseURL = "https://groupietrackers.herokuapp.com"

func FetchDates(id int, baseURL string) (models.Date, error) {
	// Corrected URL formatting
	url := fmt.Sprintf("%s/api/dates/%d", baseURL, id)

	response, err := http.Get(url)
	if err != nil {
		return models.Date{}, err // Return error instead of log.Fatal to handle errors properly
	}
	defer response.Body.Close()

	var dates models.Date
	err = json.NewDecoder(response.Body).Decode(&dates)
	if err != nil {
		return models.Date{}, err
	}
	return dates, nil
}

func DateHandlers(w http.ResponseWriter, r *http.Request) {
	datepath := "/date/"
	if r.URL.Path != datepath {
		HandlingErrors(w, "Path Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		HandlingErrors(w, "Wrong Request Method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandlingErrors(w, "Invalid date ID", http.StatusBadRequest)
		return
	}

	// Use the global baseURL when calling FetchDates
	dates, err := FetchDates(id, baseURL)
	if err != nil {
		HandlingErrors(w, "Unable to load date information", http.StatusInternalServerError)
		return
	}

	templ, err := template.ParseFiles("web/templates/layout.html", "web/templates/dates.html")
	if err != nil {
		HandlingErrors(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, dates)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		HandlingErrors(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
