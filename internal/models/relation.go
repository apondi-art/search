// internal/models/relation.go
package models

type RelationData struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

