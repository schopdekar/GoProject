// main.go
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Define a sample data structure
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Create a slice to store items (simulating a database)
// var items []Item

// // Define a handler to get all items
// func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(items)
// }

// // Define a handler to get a specific item by ID
// func GetItemHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// Get the "id" parameter from the request URL
// 	params := mux.Vars(r)
// 	itemID := params["id"]

// 	// Find the item with the given ID
// 	for _, item := range items {
// 		if item.ID == itemID {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}

// 	// If no item found, return a 404 Not Found response
// 	w.WriteHeader(http.StatusNotFound)
// 	json.NewEncoder(w).Encode(map[string]string{"error": "Item not found"})
// }

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Add routes
	// router.HandleFunc("/items", GetItemsHandler).Methods("GET")
	// router.HandleFunc("/items/{id}", GetItemHandler).Methods("GET")
	router.HandleFunc("/initiatePayment", InitiatePayment).Methods("POST")
	http.Handle("/", router)

	// Dummy data for testing
	// items = append(items, Item{ID: "1", Name: "Item 1"})
	// items = append(items, Item{ID: "2", Name: "Item 2"})

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
