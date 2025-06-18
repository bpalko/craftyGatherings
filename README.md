# Crafty Gatherings

A small business e-commerce website for showcasing and selling custom products, built with Go and a modern, responsive frontend.

## Features
- Homepage with product grid and branding
- Individual product subpages with image gallery, details, and related products
- Contact page linking to Etsy messaging
- Sticky, consistent header and responsive design
- Product data loaded from a JSON file
- SEO-friendly product URLs (e.g., `/product/groomsman-stickers`)
- Fully Etsy compliant

## Stack
- **Backend:** Go (net/http)
- **Frontend:** HTML, TailwindCSS (via CDN), custom CSS
- **Templating:** Go html/template
- **Data:** Static `products.json` file

## Repository Structure
```
craftyGatherings/
├── handlers/           # Go HTTP handlers for routing and page logic
│   └── handlers.go
├── product/            # Product struct, loader, and related logic
│   ├── product.go
│   ├── product_struct.go
├── static/             # Static assets (images, CSS)
│   ├── brandLogo.png
│   ├── styles.css
│   └── ...
├── templates/          # HTML templates for all pages
│   ├── index.html
│   ├── product.html
│   ├── contact.html
│   └── ...
├── products.json       # Product data
├── main.go             # Application entry point
├── README.md           # This file
└── ...
```

## Design Decisions
- **Go Standard Library:** Chosen for simplicity and reliability; no external web frameworks.
- **SEO-Friendly URLs:** Product pages use slugified names for clean, shareable links.
- **Product Data:** Loaded from a static JSON file for easy editing and portability.
- **Related Products:** Each product can reference related products by name, resolved to full product objects for template access.
- **Consistent Header:** All pages share a sticky, left-aligned header with custom fonts to match branding.
- **Responsive Design:** TailwindCSS and custom media queries ensure usability on all devices.
- **Contact:** No backend form processing; users are directed to Etsy for messaging, reducing spam and complexity.

## How to Run
1. Install Go (1.18+ recommended).
2. Clone this repository.
3. Run `go run main.go` from the project root.
4. Visit `http://localhost:8080` in your browser.

## Customization
- Add/edit products in `products.json`.
- Update images in the `static/` directory.
- Adjust styles in `static/styles.css` or templates as needed.

---

For questions or contributions, please open an issue or pull request.
