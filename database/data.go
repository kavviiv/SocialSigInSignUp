package database

// // FetchData :
// func FetchData() []UserData {

// 	db := OpenConn()

// 	source := UserData{}
// 	var data []UserData

// 	rows, err := db.Query("SELECT google_id, facebook_id, line_id, email FROM test")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		err := rows.Scan(&source.GoogleID, &source.FacebookID, &source.LineID, &source.Email)
// 		if err != nil {
// 			fmt.Print(err)
// 		}

// 		data = append(data, source)
// 	}

// 	return data
// }
