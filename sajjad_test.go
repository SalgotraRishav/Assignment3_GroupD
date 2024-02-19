package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// This function is created by SAJJAD KAZMI (500217679) to test the PUT api
func TestUpdateItemHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(updateItemHandler))

	defer server.Close()

payload := {"id": "123", "name": "Updated Item"}

req, err := http.NewRequest(http.MethodPut, server.URL, strings.NewReader(payload))

	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("could not send request: %v", err)
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode

	fmt.Println("The status code received is", statusCode)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status-code: 200 but received %v", resp.StatusCode)
	}
}

func updateItemHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}