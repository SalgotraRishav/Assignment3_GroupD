package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var tasks = []Task{}

const port = ":8080"

func main() {
	http.HandleFunc("/item/", taskHandler)
	fmt.Printf("Server is running on port: %v\n", port)
	http.ListenAndServe(port, nil)
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := strings.TrimPrefix(r.URL.Path, "/item/")

	switch r.Method {
	case "PUT":
		var updatedTask Task
		if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		found := false
		for i, task := range tasks {
			if task.ID == taskID {
				updatedTask.ID = task.ID
				tasks[i] = updatedTask
				found = true
				break
			}
		}
		if !found {
			http.Error(w, "The item is not found.", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(updatedTask)
	}
}
