// internal/models/combined_data.go
package models

type CombinedData struct {
	Artist Artist
	Date   Date
}

type AllLocations struct {
	Locationslist []LocationData `json:"index"`
}

type AllDates struct {
	Dateslist []Date `json:"index"`
}

type Suggestion struct {
	Value      string
	Match      string
	ArtistId   int
	ArtistName string
}
