package handlers

import (
	"craftyGatherings/product"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	if err := tmpl.Execute(w, product.Products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ProductPageHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var prod *product.Product
	for i := range product.Products {
		if product.Products[i].LocalURL == path {
			prod = &product.Products[i]
			break
		}
	}
	if prod != nil {
		tmpl := template.Must(template.ParseFiles("templates/product.html"))
		if err := tmpl.Execute(w, prod); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.NotFound(w, r)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	file := "templates/contact.html"
	if _, err := template.ParseFiles(file); err != nil {
		http.Error(w, "Contact page not found", http.StatusInternalServerError)
		return
	}
	http.ServeFile(w, r, file)
}
