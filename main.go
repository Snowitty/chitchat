package main

import (
	"log"
	"net/http"

	. "github.com/snowitty/chitchat/config"
	. "github.com/snowitty/chitchat/routes"
)

func main() {
	startWebServer()
}

func startWebServer() {
	config := LoadConfig()
	r := NewRouter()

	assets := http.FileServer(http.Dir(config.App.Static))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)

	log.Println("Starting HTTP Service at " + config.App.Address)
	err := http.ListenAndServe(config.App.Address, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + config.App.Address)
		log.Println("Error: " + err.Error())
	}
}
