package handlers

import (
	"encoding/json"
	"net/http"

	"learn.zone01kisumu.ke/git/quochieng/groupie-tracker/internal/api"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("q")
	if input == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	suggestions := api.SearchArtist(input)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(suggestions); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
