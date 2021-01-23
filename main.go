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

//แยกคอลัมน์ในการเก็บ id ของแต่ละตัว
//ทำ register บันทึก id + email to db

// ต้องเก็บ email ของ facebook ไหมหรือเก็บแค่ของ google
// ถ้าเก็บต้องแปลง u0040 เป็น @ ไหม และจะเอาไปใช้ไหม
// ถ้าแปลงมีวิธียังไง
// ถ้าไม่เก็บ
