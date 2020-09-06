package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "43880"
	if len(os.Args) > 1 {
		fmt.Println(len(os.Args))
		port = string(os.Args[1])
	}
	http.Handle("/", http.FileServer(assets))
	log.Print("Started HTTP server on " + ":" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server: ", err)
		return
	}
}
