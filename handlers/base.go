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

	// // Root
	mux.Handle("/", http.FileServer(http.Dir("templates/")))
	mux.HandleFunc("/facebook/login", handleFacebookLogin)
	mux.HandleFunc("/facebookcallback", handleFacebookCallback)

	for _, el := range fetch {
		// if el.GoogleID != nil {
		// mux.HandleFunc("/google/login", handleGoogleLogin)
		// mux.HandleFunc("/googlecallback", handleGoogleCallback)
		// fmt.Println("Google Login")
		// }

		if el.GoogleID == nil {
			mux.HandleFunc("/google/register", handleGoogleRegister)
			mux.HandleFunc("/googlecallback", handleGoogleRegisterCallback)
			fmt.Println("Google ID is nill")
		}

		// if el.FacebookID != nil {
		// 	mux.HandleFunc("/facebook/login", handleFacebookLogin)
		// 	mux.HandleFunc("/facebookcallback", handleFacebookCallback)
		// 	fmt.Println("Facebook Login")
		// }

		// if el.FacebookID == nil {
		// 	mux.HandleFunc("/facebookcallback", handleFacebookRegister)
		// 	fmt.Println("Facebook ID is nill")
		// }

		// if el.LineID != nil {
		// 	mux.HandleFunc("/line/login", handleLineLogin)
		// 	mux.HandleFunc("/linecallback", handleLineCallback)
		// 	fmt.Println("Line Login")
		// }

		// if el.LineID == nil {
		// 	mux.HandleFunc("/linecallback", handleLineRegister)
		// 	fmt.Println("Line ID is nill")
		// }
	}
	return mux
}
