package product

// Product represents a single product with its details.
type Product struct {
	ID          string   `json:"id"`
	Name        string   `json:"Name"`
	Images      []string `json:"images"`
	Description string   `json:"description,omitempty"` // Optional field for product description
	EtsyURL     string   `json:"etsy_url"`
	Type        string   `json:"type"`
	Price       string   `json:"price"`
	Reviews     []string `json:"reviews,omitempty"`
	LocalURL    string   `json:"local_url"`
	Related     *Product `json:"-"` // Related products as pointers
}
