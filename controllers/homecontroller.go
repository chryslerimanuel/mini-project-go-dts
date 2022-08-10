package controllers

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Load template index html
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Printf("Errors %s load template", err)
		http.Error(w, "An error has occured.", http.StatusInternalServerError)
		return
	}

	// Execute template index html
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Errors %s execute template", err)
		http.Error(w, "An error has occured.", http.StatusInternalServerError)
		return
	}
}
