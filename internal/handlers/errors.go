package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type ErrorData struct {
	ErrorMessage string
	StatusCode   int
}
// HandlingErrors sends a custom error response to the client using a template.
// It takes in the response writer, an error message, a status code,
// and an optional templateLoader function.
// The function will render a template for the error page, 
//or fall back to default templates if none are provided.

func HandlingErrors(w http.ResponseWriter, errorMessage string, statuscode int, templateLoader ...func() (*template.Template, error)) {
	w.WriteHeader(statuscode)
	errorData := ErrorData{
		ErrorMessage: errorMessage,
		StatusCode:   statuscode,
	}

	var temp *template.Template
	var err error

	// If a templateLoader function is provided, use it; otherwise, load the default templates
	if len(templateLoader) > 0 {
		temp, err = templateLoader[0]()
	} else {
		temp, err = template.ParseFiles("web/templates/layout.html", "web/templates/error.html")
	}
	if err != nil {
		// Log and return an internal server error if template parsing fails.
		log.Printf("Error parsing error template: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, errorData)
	if err != nil {
		// Log and return an internal server error if template execution fails.
		log.Printf("Error executing error template: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
