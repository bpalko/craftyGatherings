package product

import (
	"craftyGatherings/util"
	"encoding/json"
	"log"
	"math/rand"
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

	// Set LocalURL for each product
	for i := range Products {
		Products[i].LocalURL = "/product/" + util.Slugify(Products[i].Name)
	}

	// Grab a random product from the list of Products and store in Related field as a pointer,
	// ensuring it's not the same product.
	for i := range Products {
		if len(Products) > 1 {
			// Build a list of indices excluding the current product
			indices := make([]int, 0, len(Products)-1)
			for j := range Products {
				if j != i {
					indices = append(indices, j)
				}
			}
			// Pick a random index from the list
			randomIdx := indices[rand.Intn(len(indices))]
			Products[i].Related = &Products[randomIdx]
		} else {
			Products[i].Related = nil
		}
	}
}
