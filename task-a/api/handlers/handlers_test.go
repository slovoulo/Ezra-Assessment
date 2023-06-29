package handlers

import (
	
	
	"net/http"
	"net/http/httptest"
	"testing"
	
)

func TestHomeHandler(t *testing.T) {
	// Create a sample request
	req, err := http.NewRequest("GET", "/v1/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	responseRecorder := httptest.NewRecorder()

	// Call the HomeHandler function
	HomeHandler(responseRecorder, req)

	// Check the response status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, responseRecorder.Code)
	}

	// Check the response body
	expectedBody := "Welcome to the Elevator app!"
	if responseRecorder.Body.String() != expectedBody {
		t.Errorf("expected body '%s' but got '%s'", expectedBody, responseRecorder.Body.String())
	}
}

