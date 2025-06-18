package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Response represents the structure of the Etsy API response
type Response struct {
	Results []Product `json:"results"`
}

// FetchEtsyProducts fetches product data from Etsy's API
func FetchEtsyProducts(apiKey string) ([]Product, error) {
	url := fmt.Sprintf("https://openapi.etsy.com/v2/listings/active?api_key=%s", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from Etsy: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return response.Results, nil
}

// WriteToJSONFile writes the product data to a local JSON file
func WriteToJSONFile(filename string, data []Product) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if err := ioutil.WriteFile(filename, file, 0644); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
