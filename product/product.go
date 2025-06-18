package product

import (
	"encoding/json"
	"log"
	"os"
)

var Products []Product

// LoadProducts loads product data from the JSON file into the global Products variable.
func LoadProducts() {
	file, err := os.Open("products.json")
	if err != nil {
		log.Fatalf("Error loading products: %v", err)
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&Products); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
}
