package main

import (
	"log"
	"net/http"

	"github.com/0xEDU/goblog/cmd/ui/http_handlers"
)

const (
	GrpcPort   = "9090"
	ServerPort = "8080"
)

func main() {
	http.HandleFunc("/", http_handlers.HomeHandler)
	http.HandleFunc("/css/", http_handlers.CssHandler)
	log.Println("Server running on port " + ServerPort)
	http.ListenAndServe(":"+ServerPort, nil)
}
