package handlers

import (
	"net/http"

	"github.com/snowitty/chitchat/models"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	threads, err := models.Threads()
	if err == nil {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "auth.navbar", "index")
		}
	}
}
