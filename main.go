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
	// log.Fatalln(server)
	// _ = server

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	log.Println()
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}

	// fmt.Println()
	// fmt.Println("Starting HTTP Server. Listening at 127.0.0.1:9090")

	// data := database.CheckData()
	// fmt.Println(data)
	// fmt.Println("Main")
	// // http.HandleFunc("/", handleMain)

	// data := database.CheckData()
	// fmt.Println("Data =", data)
	// for _, el := range data {
	// 	fmt.Println("Nill")
	// 	if el.GoogleID == nil {
	// 		fmt.Println("nill")
	// 		// // 		fmt.Println()
	// 		// // 		fmt.Println("Function handler main")
	// 		// // 		log.Fatalln("Handler Main")
	// 		// // 		// http.HandleFunc("/google/login", handleGoogleLogin)
	// 		// // 		// http.HandleFunc("/googlecallback", handleGoogleCallback)
	// 	}
	// }

	// http.Handle("/", http.FileServer(http.Dir("templates/")))
	// http.HandleFunc("/google/login", handlers.handleGoogleLogin)

	// fmt.Println(http.ListenAndServe("127.0.0.1:9090", nil))
}

// func handleMain(w http.ResponseWriter, r *http.Request) {
// 	data := database.CheckData()
// 	fmt.Println("Data =", data)
// 	for _, el := range data {
// 		if el.GoogleID == "" {
// 			fmt.Println()
// 			// fmt.Println("Function handler main")
// 			// log.Fatalln("Handler Main")
// 			http.HandleFunc("/google/login", handlers.handleGoogleLogin())
// 			http.HandleFunc("/googlecallback", handleGoogleCallback)
// 		}
// 	}
// 	return
// }

//แยกคอลัมน์ในการเก็บ id ของแต่ละตัว
//ทำ register บันทึก id + email to db

// ต้องเก็บ email ของ facebook ไหมหรือเก็บแค่ของ google
// ถ้าเก็บต้องแปลง u0040 เป็น @ ไหม และจะเอาไปใช้ไหม
// ถ้าแปลงมีวิธียังไง
// ถ้าไม่เก็บ
