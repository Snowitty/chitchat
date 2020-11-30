package handlers

import (
	"html/template"
	"net/http"

	"github.com/snowitty/chitchat/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	files := []string{"views/layout.html", "views/navbar.html", "views/index.html"}
	templates := template.Must(template.Must(template.ParseFiles(files...)))
	threads, err := models.Threads()
	if err != nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
