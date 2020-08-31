package main

import (
	"log"
	"net/http"
	"os"
)

const (
	// Host name of the HTTP Server
	Host = "localhost"
)

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/index.html", http.StatusSeeOther)
	})

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	err := http.ListenAndServe(Host+":"+string(os.Args[1]), nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server: ", err)
		return
	}
}
