package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"learn.zone01kisumu.ke/git/quochieng/groupie-tracker/internal/models"
)

func TestFetchDates(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Define the mock response
		dates := models.Date{
			ID:    1,
			Dates: []string{"2024-01-01", "2024-01-02"},
		}
		// Set content type and write JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dates)
	}))
	defer server.Close()

	// Call the FetchDates function with the mock server's URL
	id := 1
	got, err := FetchDates(id, server.URL)
	if err != nil {
		t.Fatalf("FetchDates() returned error: %v", err)
	}

	// Define the expected result
	want := models.Date{
		ID:    1,
		Dates: []string{"2024-01-01", "2024-01-02"},
	}

	// Check if the result matches the expected result
	if got.ID != want.ID || len(got.Dates) != len(want.Dates) {
		t.Errorf("Expected %v, got %v", want, got)
	}

	for i := range got.Dates {
		if got.Dates[i] != want.Dates[i] {
			t.Errorf("Expected %v, got %v", want.Dates[i], got.Dates[i])
		}
	}
}
