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

func serveFiles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", serveFiles)
	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server: ", err)
		return
	}
}
