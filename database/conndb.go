package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// OpenConn :
func OpenConn() *sql.DB {
	conndb := "user=postgres dbname=postgres password=130242 host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", conndb)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("Successfully connected!")
	fmt.Println()

	return db
}

// // InsertData :
// func InsertData() {
// 	db := OpenConn()
// 	defer db.Close()

// }

// FetchData :
func FetchData() []UserData {
	db := OpenConn()
	defer db.Close()

	source := UserData{}
	var data []UserData

	// WHERE return user_id
	rows, err := db.Query("SELECT google_id, facebook_id, line_id, email FROM test")
	if err != nil {
		log.Fatalln("Error database query in Fetch Data")
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&source.GoogleID, &source.FacebookID, &source.LineID, &source.Email)
		if err != nil {
			fmt.Println("Error row scan in Fetch Data")
			fmt.Print(err)
		}
		data = append(data, source)
	}
	return data
}

// CheckData :
func CheckData() []UserData {
	db := OpenConn()
	defer db.Close()

	source := UserData{}
	var data []UserData

	// WHERE return user_id
	rows, err := db.Query("SELECT google_id, facebook_id, line_id, email FROM test WHERE user_id='654321'")
	if err != nil {
		log.Fatalln("Error database query in Check Data")
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&source.GoogleID, &source.FacebookID, &source.LineID, &source.Email)
		if err != nil {
			fmt.Println(("Error row scan in Check Data"))
			fmt.Println(err)
			// data = append(data, source)
		}

		data = append(data, source)
	}
	fmt.Println()
	return data
}
