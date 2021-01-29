package main

import (
	"Project/login/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    fmt.Sprint("127.0.0.1:9090"),
		Handler: handlers.New(),
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	log.Println()
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
