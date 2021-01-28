package handlers

import (
	"Project/login/database"
	"fmt"
	"net/http"
)

// New :
func New() http.Handler {
	mux := http.NewServeMux()

	fetch := database.FetchData()
	// fmt.Println("Data Fetch =", fetch)

	// check := database.CheckData()
	// fmt.Println("Data Check =", check)

	// Root
	mux.Handle("/", http.FileServer(http.Dir("templates/")))
	// mux.HandleFunc("/google/login", handleGoogleLogin)
	// mux.HandleFunc("/googlecallback", handleGoogleCallback)
	mux.HandleFunc("/facebook/login", handleFacebookLogin)
	mux.HandleFunc("/facebookcallback", handleFacebookCallback)
	mux.HandleFunc("/line/login", handleLineLogin)
	mux.HandleFunc("/linecallback", handleLineCallback)

	for _, el := range fetch {
		if el.GoogleID != nil {
			mux.HandleFunc("/google/login", handleGoogleLogin)
			mux.HandleFunc("/googlecallback", handleGoogleCallback)
		}

		if el.GoogleID == nil {
			// mux.HandleFunc("/googlecallback", handleGoogleRegister)
			fmt.Println("Nill")
		}
	}

	// mux.HandleFunc("/googlecallback", handleGoogleRegister)

	// datafetch := database.FetchData()
	// fmt.Println("Data Fetch =", datafetch)

	// check := database.CheckData()
	// fmt.Println("Data Check =", check)

	// for _, el := range check {
	// 	if el.GoogleID == nil {
	// mux.HandleFunc("/googlecallback", handleGoogleRegister)
	// 	}
	// }

	// if data != nil {
	// 	mux.HandleFunc("/google/login", handleGoogleLogin)
	// 	mux.HandleFunc("/googlecallback", handleGoogleCallback)
	// 	mux.HandleFunc("/facebook/login", handleFacebookLogin)
	// 	mux.HandleFunc("/facebookcallback", handleFacebookCallback)
	// 	mux.HandleFunc("/facebook/regist", handleFacebookRegister)
	// 	mux.HandleFunc("/line/login", handleLineLogin)
	// 	mux.HandleFunc("/linecallback", handleLineCallback)
	// }

	// if data == nil {
	// 	mux.HandleFunc("/googlecallback", handleGoogleRegister)
	// }

	// mux.HandleFunc("/line/regist", handleGoogleLogin)
	// mux.HandleFunc("/google/regist", handleGoogleLogin)

	return mux
}
