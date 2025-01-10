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
// FetchLocationData fetches location data for a specific artist using the provided artist ID.
// It sends a GET request to the appropriate API endpoint, decodes the JSON response, 
// and returns the models.LocationData object or an error if something goes wrong.

func FetchLocationData(id int) (models.LocationData, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", id)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var locations models.LocationData
	err = json.NewDecoder(response.Body).Decode(&locations)
	if err != nil {
		return models.LocationData{}, err
	}
	return locations, nil
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	pathlocation := "/location/"
	if r.URL.Path != pathlocation {
		HandlingErrors(w, "Path Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		HandlingErrors(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandlingErrors(w, "Invalid location ID", http.StatusBadRequest)
		return
	}

	location, err := FetchLocationData(id)
	if err != nil {
		HandlingErrors(w, "Unable to load location data", http.StatusInternalServerError)
		return
	}

	templ, err := template.ParseFiles("web/templates/layout.html", "web/templates/location.html")
	if err != nil {
		HandlingErrors(w, "Error Passing Template File", http.StatusInternalServerError)
	}

	err = templ.Execute(w, location)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		HandlingErrors(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
