package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"learn.zone01kisumu.ke/git/quochieng/groupie-tracker/internal/models"
)

func FetchRelation(client *http.Client, id int) (models.RelationData, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id)
	response, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	var relation models.RelationData
	err = json.NewDecoder(response.Body).Decode(&relation)
	if err != nil {
		return models.RelationData{}, err
	}
	return relation, nil
}

func RelationHandler(w http.ResponseWriter, r *http.Request) {
	pathrelation := "/relation/"
	if r.URL.Path != pathrelation {
		HandlingErrors(w, "Path Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		HandlingErrors(w, "Method Not allowed", http.StatusMethodNotAllowed)
		return
	}
	relationid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(relationid)
	if err != nil {
		log.Printf("Error converting relation  ID: %v", err)
		HandlingErrors(w, "Invalid relation ID", http.StatusBadRequest)
		return
	}
	client := &http.Client{}
	relation, err := FetchRelation(client, id)
	if err != nil {
		HandlingErrors(w, "Unable to load realtion  data", http.StatusInternalServerError)
		return
	}
	temp, err := template.ParseFiles("web/templates/layout.html", "web/templates/relation.html")
	if err != nil {
		HandlingErrors(w, "Error loading template file", http.StatusNotFound)
	}
	err = temp.Execute(w, relation)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		HandlingErrors(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
