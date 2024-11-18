package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Model
type Item struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"unique"`
	Value int    `json:"value"`
}

func main() {
	// Initialize DB
	var err error
	// Use environment variables for database configuration
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&Item{})

	// Generate and insert/update 1000 random data
	generateData()

	// Initialize Router
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	// Start Server
	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Generate and insert/update 1000 random data
func generateData() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		name := fmt.Sprintf("Item %d", i+1)
		value := rand.Intn(1000)

		var item Item
		if err := db.Where("name = ?", name).First(&item).Error; err == nil {
			// Update the existing item
			item.Value = value
			db.Save(&item)
		} else {
			// Create a new item
			item = Item{Name: name, Value: value}
			db.Create(&item)
		}
	}
}

// Handlers

// Get all items
func getItems(w http.ResponseWriter, r *http.Request) {
	var items []Item
	db.Find(&items)
	json.NewEncoder(w).Encode(items)
}

// Get a single item by ID
func getItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item Item
	if err := db.First(&item, params["id"]).Error; err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

// Create a new item
func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := db.Create(&item).Error; err != nil {
		http.Error(w, "Could not create item", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

// Update an item by ID
func updateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item Item
	if err := db.First(&item, params["id"]).Error; err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	db.Save(&item)
	json.NewEncoder(w).Encode(item)
}

// Delete an item by ID
func deleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item Item
	if err := db.First(&item, params["id"]).Error; err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	db.Delete(&item)
	w.WriteHeader(http.StatusNoContent)
}
