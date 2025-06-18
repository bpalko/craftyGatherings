package main

import (
	"craftyGatherings/handlers"
	"craftyGatherings/product"
	"log"
	"net/http"
)

func main() {
	// apiKey := "your_etsy_api_key" // Replace with your Etsy API key
	// filename := "products.json"

	// // Fetch products from Etsy
	// products, err := product.FetchEtsyProducts(apiKey)
	// if err != nil {
	// 	log.Fatalf("Error fetching products: %v", err)
	// }

	// Write products to a JSON file
	// if err := product.WriteToJSONFile(filename, products); err != nil {
	// 	log.Fatalf("Error writing to JSON file: %v", err)
	// }
	// log.Printf("Successfully wrote %d products to %s", len(products), filename)

	// Load products from the JSON file
	product.LoadProducts()

	// Start the server
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)
	http.HandleFunc("/product", handlers.RedirectHandler)
	http.HandleFunc("/product/details", handlers.ProductPageHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
