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
		Handler: handlers.Oauth(),
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	log.Println()
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}

	// paths := []string{"/google/login", "/googlecallback", "/facebook/login", "/facebookcallback", "/line/login", "linecallback"}
	// for i := 0; i < len(paths); i++ {
	// 	fmt.Println(paths[2])
	// 	break
	// }
	// for i, path := range paths {
	// 	fmt.Println(i, path[i])
	// }
}
