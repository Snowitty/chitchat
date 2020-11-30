package handlers

import (
	"net/http"

	"github.com/snowitty/chitchat/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := models.Threads()
	if err == nil {
		generateHTML(w, threads, "layout", "navbar", "index")
	}
}
