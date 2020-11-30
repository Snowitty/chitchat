package main

import (
	"log"
	"net/http"

	. "github.com/snowitty/chitchat/routes"
)

func main() {

}

func startWebServer(port string) {
	r := NewRouter()
	http.Handle("/", r)

	log.Println("Starting HTTP Service at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
