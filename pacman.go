package main

import (
	"log"
	"net/http"
)

const (
	// Host name of the HTTP Server
	Host = "localhost"
	// Port for HTTP server
	Port = "8080"
)

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/index.html", http.StatusSeeOther)
	})

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server: ", err)
		return
	}
}
