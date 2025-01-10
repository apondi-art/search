package handlers

import (
	"html/template"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlingErrors(t *testing.T) {
	// Mock template loader for testing
	mockLoader := func() (*template.Template, error) {
		return template.New("test").Parse("Error: {{.ErrorMessage}}, Status: {{.StatusCode}}")
	}

	// Define test cases
	testcases := []struct {
		Name            string
		ErrorMessage    string
		ErrorStatusCode int
		ExpectedMessage string
		ExpectedCode    int
	}{
		{
			Name:            "internalserver",
			ErrorMessage:    "Internal Server Error",
			ErrorStatusCode: 500,
			ExpectedMessage: "Internal Server Error",
			ExpectedCode:    500,
		},
		{
			Name:            "pathNotfound", // fixed typo "pathNotforund"
			ErrorMessage:    "Path Not Found",
			ErrorStatusCode: 404,
			ExpectedMessage: "Path Not Found",
			ExpectedCode:    404,
		},
		{
			Name:            "methodNotAllowed",
			ErrorMessage:    "Method Not Allowed",
			ErrorStatusCode: 405,
			ExpectedMessage: "Method Not Allowed", 
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			// Create a new ResponseRecorder (which implements http.ResponseWriter)
			wr := httptest.NewRecorder()

			// Call the function being tested, injecting the mock template loader
			HandlingErrors(wr, tc.ErrorMessage, tc.ErrorStatusCode, mockLoader)

			// Check the status code
			if wr.Code != tc.ErrorStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					wr.Code, tc.ErrorStatusCode)
			}

			// Check the response body content
			if !strings.Contains(wr.Body.String(), tc.ExpectedMessage) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					wr.Body.String(), tc.ExpectedMessage)
			}
		})
	}
}
