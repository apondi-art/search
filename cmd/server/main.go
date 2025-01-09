// cmd/server/main.go
package main

import (
	"log"
	"net/http"

	"learn.zone01kisumu.ke/git/quochieng/groupie-tracker/internal/handlers"
)

// main connects other endpoints such as location,artist,date and relation
func main() {

	http.Handle("/static/", handlers.CustomHandler("./web/static"))
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artist/", handlers.ArtistsHandler)
	http.HandleFunc("/location/", handlers.LocationHandler)
	http.HandleFunc("/date/", handlers.DateHandlers)
	http.HandleFunc("/relation/",handlers.RelationHandler)
	http.HandleFunc("/search", handlers.SearchHandler)

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
