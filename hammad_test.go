package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Hammad Ul Hassan 500230292 - Code for testing the ItemsHandler function
func TestCreateItem(t *testing.T) {
	// Create a new HTTP POST request with a Item JSON body
	var jsonStr = []byte(`{"title":"Test Item", "status":"pending"}`)
	req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(itemsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var item Item
	if err := json.NewDecoder(rr.Body).Decode(&item); err != nil {
		t.Fatalf("Could not decode response body: %v", err)
	}

	// Validate the Item fields
	if item.Title != "Test Item" {
		t.Errorf("handler returned unexpected body for title: got %v want %v", item.Title, "Test Item")
	}
	if item.Status != "pending" {
		t.Errorf("handler returned unexpected body for status: got %v want %v", item.Status, "pending")
	}
	if item.ID == "" {
		t.Errorf("handler did not generate an ID for the item")
	}
}

