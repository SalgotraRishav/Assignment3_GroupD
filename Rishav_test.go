package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// This function is created by Rishav (500228178) to test the GET api
func TestItemsHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(itemsHandler))
	resp, err := http.Get(server.URL)
	statusCode := resp.StatusCode

	// Printing messages on the terminal
	fmt.Println("The status code received is ", statusCode)
	fmt.Println("The value of err is ", err)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status-code: 200 but received %v", resp.StatusCode)
	}
}
