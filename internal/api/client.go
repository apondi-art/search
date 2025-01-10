// internal/api/client.go
package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"learn.zone01kisumu.ke/git/quochieng/groupie-tracker/internal/models"
	// "groupie-tracker/internal/models"
)

var (
	ArtistsURL  = "https://groupietrackers.herokuapp.com/api/artists"
	LocationUrl = "https://groupietrackers.herokuapp.com/api/locations"
	DateUrl     = "https://groupietrackers.herokuapp.com/api/dates"
	RelationUrl = "https://groupietrackers.herokuapp.com/api/relation"
	artists     []models.Artist
	locations   models.AllLocations
	dates       models.AllDates
)

func init() {
	FetchAllArtist()
	FetchAllartistlocations()
	FetchAllAristDates()
}

func FetchJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func FetchAllArtist() {
	if err := FetchJSON(ArtistsURL, &artists); err != nil {
		log.Fatal("Error fetching artist data", err)
	}

}

func FetchAllartistlocations() {
	if err := FetchJSON(LocationUrl, &locations); err != nil {
		log.Fatal(err)
	}

}
func FetchAllAristDates() {
	if err := FetchJSON(DateUrl, &dates); err != nil {
		log.Fatal(err)
	}

}

func FetchAllData() ([]models.CombinedData, error) {
	var artists []models.Artist

	if err := FetchJSON(ArtistsURL, &artists); err != nil {
		return nil, err
	}

	combinedData := make([]models.CombinedData, len(artists))
	for i, artist := range artists {
		combinedData[i] = models.CombinedData{
			Artist: artist,
		}
	}

	return combinedData, nil
}

func Suggestn(values string, match string, id int, name string) models.Suggestion {
	return models.Suggestion{
		Value:      values,
		Match:      match,
		ArtistId:   id,
		ArtistName: name,
	}
}

func SearchArtist(input interface{}) []models.Suggestion {
	inputStr := strings.ToLower(fmt.Sprintf("%v", input))
	resultsChan := make(chan []models.Suggestion, len(artists))
	var wg sync.WaitGroup

	if len(artists) == 0 {
		fmt.Println("No artists found")
		return nil
	}

	// Create maps to associate locations and dates with artist IDs
	locationMap := make(map[int][]string)
	dateMap := make(map[int][]string)

	// Pre-process locations and dates to map them to artist IDs
	for i, loc := range locations.Locationslist {
		if i < len(artists) { // Ensure we don't go out of bounds
			locationMap[artists[i].ID] = loc.Locations
		}
	}

	for i, date := range dates.Dateslist {
		if i < len(artists) { // Ensure we don't go out of bounds
			dateMap[artists[i].ID] = date.Dates
		}
	}

	for _, artist := range artists {
		wg.Add(1)
		go func(artist models.Artist) {
			defer wg.Done()
			var suggestions []models.Suggestion

			// Artist name search
			if strings.Contains(strings.ToLower(artist.Name), inputStr) {
				suggestions = append(suggestions, Suggestn("ArtistBand", artist.Name, artist.ID, ""))
			}

			// First album search
			if strings.Contains(strings.ToLower(artist.FirstAlbum), inputStr) {
				suggestions = append(suggestions, Suggestn("FirstAlbum", artist.FirstAlbum, artist.ID, artist.Name))
			}

			// Creation date search
			if strings.Contains(strconv.Itoa(artist.CreationDate), inputStr) {
				suggestions = append(suggestions, Suggestn("CreationDate", strconv.Itoa(artist.CreationDate), artist.ID, artist.Name))
			}

			// Members search
			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), inputStr) {
					suggestions = append(suggestions, Suggestn("Member", member, artist.ID, artist.Name))
				}
			}

			// Locations search using the map
			if artistLocations, ok := locationMap[artist.ID]; ok {
				for _, loc := range artistLocations {
					if strings.Contains(strings.ToLower(loc), inputStr) {
						suggestions = append(suggestions, Suggestn("Location", loc, artist.ID, artist.Name))
					}
				}
			}

			// Dates search using the map
			if artistDates, ok := dateMap[artist.ID]; ok {
				for _, date := range artistDates {
					if strings.Contains(strings.ToLower(date), inputStr) {
						suggestions = append(suggestions, Suggestn("Concert Date", date, artist.ID, artist.Name))
					}
				}
			}

			resultsChan <- suggestions
		}(artist)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var allSuggestions []models.Suggestion
	for suggestions := range resultsChan {
		allSuggestions = append(allSuggestions, suggestions...)
	}

	return allSuggestions
}
