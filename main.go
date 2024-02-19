// Start of Ameen's code.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// Item represents a to-do item.
type Item struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"` // e.g., "pending" or "completed"
}

var items = []Item{} // In-memory storage for items
const Dport = ":8012"

func main() {
	http.HandleFunc("/items", itemsHandler)
	http.HandleFunc("/item/", itemHandler)
	fmt.Printf("Server is starting on port: %v\n", Dport) // Added newline for better terminal output
	http.ListenAndServe(Dport, nil)
}

// End of Ameen's code.

// Handle requests to the /items endpoint
func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// Start of Ramanpreet's code
	case "GET":
		json.NewEncoder(w).Encode(items)
	// End of Ramanpreet's code

	// Start of Anas Basheer's code
	case "POST":
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item.ID = uuid.New().String() // Generate a unique ID for the item
		items = append(items, item)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(item)
	// End of Anas Basheer's code

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Handle requests to the /item/{id} endpoint
func itemHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the item ID from the URL path
	itemID := strings.TrimPrefix(r.URL.Path, "/item/")

	switch r.Method {

	// Start of Kiranjeet's code
	case "PUT":
		var updatedItem Item
		if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		found := false
		for i, item := range items {
			if item.ID == itemID {
				updatedItem.ID = item.ID // Ensure the ID remains unchanged
				items[i] = updatedItem
				found = true
				break
			}
		}
		if !found {
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(updatedItem)
	// End of Kiranjeet's code

	// Start of Prateek's code
	case "DELETE":
		index := -1
		for i, item := range items {
			if item.ID == itemID {
				index = i
				break
			}
		}
		if index != -1 {
			items = append(items[:index], items[index+1:]...)
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Item not found", http.StatusNotFound)
		}
	}
	// End of Prateek's code
}
