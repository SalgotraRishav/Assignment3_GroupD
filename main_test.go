package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteItemHandler(t *testing.T) {
    // Reset the items slice and add a test item
    items = []Item{{ID: "test1", Title: "Test Item"}}

    // Create a DELETE request for the test item
    req, err := http.NewRequest("DELETE", "/delete/item?id=test1", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Record the response using httptest
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(itemHandler)

    // Give the DELETE request to the handler
    handler.ServeHTTP(rr, req)

    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Decode the JSON response
    var got string
    err = json.Unmarshal(rr.Body.Bytes(), &got)
    if err != nil {
        t.Fatalf("could not read the response: %v", err)
    }

    // Define the expected response
    expected := "Item deleted successfully"
    if got != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", got, expected)
    }
}
