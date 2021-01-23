package database

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// // var db *sql.DB

// // UserData :
// type UserData struct {
// 	GoogleID   string `json:"googleID"`
// 	FacebookID string `json:"facebookID"`
// 	LineID     string `json:"lineID"`
// 	Email      string `json:"email"`
// }

// // ConnDb :
// // func ConnDb(db string) {
// // 	conndb := "user=postgres dbname=postgres password=293161 host=127.0.0.1 sslmode=disable"
// // 	db, err := sql.Open("postgres", conndb)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	// defer db.Close()
// // 	return
// // }

// func init(){

// }

// // FetchData :
// func FetchData() []UserData {
// 	conndb := "user=postgres dbname=postgres password=293161 host=127.0.0.1 sslmode=disable"
// 	db, err := sql.Open("postgres", conndb)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	source := UserData{}
// 	var data []UserData

// 	rows, err := db.Query("SELECT google_id, facebook_id, line_id, email FROM test")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 		err := rows.Scan(&source.GoogleID, &source.FacebookID, &source.LineID, &source.Email)
// 		if err != nil {
// 			fmt.Print(err)
// 		}

// 		data = append(data, source)
// 	}
// 	defer rows.Close()
// 	return data
// }
