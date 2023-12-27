package main

import (
	"net/http"
	"rmccallister/pkg/handlers"
	"rmccallister/pkg/jsonHandlers"
)

//port number

const portNumber = ":3117"

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", jsonHandlers.ServeJson)
	println("Starting application on port", portNumber)
	http.ListenAndServe(portNumber, nil)
}
