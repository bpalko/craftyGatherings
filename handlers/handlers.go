package handlers

import (
	"craftyGatherings/product"
	"craftyGatherings/util"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	if err := tmpl.Execute(w, product.Products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if i := util.Atoi(id); i >= 0 && i < len(product.Products) {
		http.Redirect(w, r, product.Products[i].EtsyURL, http.StatusFound)
		return
	}
	http.NotFound(w, r)
}

func ProductPageHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if i := util.Atoi(id); i >= 0 && i < len(product.Products) {
		tmpl := template.Must(template.ParseFiles("templates/product.html"))
		if err := tmpl.Execute(w, product.Products[i]); err != nil {
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
