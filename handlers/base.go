package handlers

import (
	"net/http"
)

// Oauth :
func Oauth() http.Handler {
	mux := http.NewServeMux()
	// fetch := database.FetchData()

	// // Root
	mux.Handle("/", http.FileServer(http.Dir("templates/")))
	mux.HandleFunc("/google/login", handleGoogleLogin)
	mux.HandleFunc("/googlecallback", handleGoogleCallback)
	mux.HandleFunc("/facebook/login", handleFacebookLogin)
	mux.HandleFunc("/facebookcallback", handleFacebookCallback)
	mux.HandleFunc("/line/login", handleLineLogin)
	mux.HandleFunc("/linecallback", handleLineCallback)

	// for _, el := range fetch {
	// 	if el.GoogleID != nil && el.Email != nil && el.FacebookID != nil || el.LineID != nil {
	// 		login()
	// 	}
	// 	// if el.GoogleID != nil {
	// 	// mux.HandleFunc("/google/login", handleGoogleLogin)
	// 	// mux.HandleFunc("/googlecallback", handleGoogleCallback)
	// 	// fmt.Println("Google Login")
	// 	// }

	// if el.GoogleID == nil {
	// 	register()
	// 	// 		mux.HandleFunc("/google/register", handleGoogleRegister)
	// 	// 		mux.HandleFunc("/googlecallback", handleGoogleRegisterCallback)
	// 	// 		fmt.Println("Google ID is nill")
	// }

	// 	// if el.FacebookID != nil {
	// 	// 	mux.HandleFunc("/facebook/login", handleFacebookLogin)
	// 	// 	mux.HandleFunc("/facebookcallback", handleFacebookCallback)
	// 	// 	fmt.Println("Facebook Login")
	// 	// }

	// 	// if el.FacebookID == nil {
	// 	// 	mux.HandleFunc("/facebookcallback", handleFacebookRegister)
	// 	// 	fmt.Println("Facebook ID is nill")
	// 	// }

	// 	// if el.LineID != nil {
	// 	// 	mux.HandleFunc("/line/login", handleLineLogin)
	// 	// 	mux.HandleFunc("/linecallback", handleLineCallback)
	// 	// 	fmt.Println("Line Login")
	// 	// }

	// 	if el.LineID == nil {
	// 		mux.HandleFunc("/linecallback", handleLineRegister)
	// 		fmt.Println("Line ID is nill")
	// 	}
	// }
	return mux
}

// Login :
func login() {
	// data := database.FetchData()
	http.HandleFunc("/google/login", handleGoogleLogin)
	http.HandleFunc("/googlecallback", handleGoogleCallback)
	http.HandleFunc("/facebook/login", handleFacebookLogin)
	http.HandleFunc("/facebookcallback", handleFacebookCallback)
	http.HandleFunc("/line/login", handleLineLogin)
	http.HandleFunc("/linecallback", handleLineCallback)
	// register()
	return
}

// Register :
func register() {
	// mux := http.NewServeMux()
	// fetch := database.FetchData()

	// // Root
	// mux.Handle("/", http.FileServer(http.Dir("templates/")))
	// mux.HandleFunc("/google/register", handleGoogleRegister)
	// mux.HandleFunc("/googlecallback", handleGoogleRegisterCallback)
	// mux.HandleFunc("/facebook/login", handleFacebookRegister)
	// mux.HandleFunc("/facebookcallback", handleFacebookCallback)
	// mux.HandleFunc("/line/login", handleLineLogin)
	// mux.HandleFunc("/linecallback", handleLineCallback)

	// for _, el := range fetch {
	// 	// if el.GoogleID != nil {
	// 	// mux.HandleFunc("/google/login", handleGoogleLogin)
	// 	// mux.HandleFunc("/googlecallback", handleGoogleCallback)
	// 	// fmt.Println("Google Login")
	// 	// }

	// if el.GoogleID == nil {
	http.HandleFunc("/google/register", handleGoogleRegister)
	http.HandleFunc("/googlecallback", handleGoogleRegisterCallback)
	// 		fmt.Println("Google ID is nill")
	// }

	// 	// if el.FacebookID != nil {
	// 	// 	mux.HandleFunc("/facebook/login", handleFacebookLogin)
	// 	// 	mux.HandleFunc("/facebookcallback", handleFacebookCallback)
	// 	// 	fmt.Println("Facebook Login")
	// 	// }

	// 	// if el.FacebookID == nil {
	// 	// 	mux.HandleFunc("/facebookcallback", handleFacebookRegister)
	// 	// 	fmt.Println("Facebook ID is nill")
	// 	// }

	// 	// if el.LineID != nil {
	// 	// 	mux.HandleFunc("/line/login", handleLineLogin)
	// 	// 	mux.HandleFunc("/linecallback", handleLineCallback)
	// 	// 	fmt.Println("Line Login")
	// 	// }

	// 	if el.LineID == nil {
	// 		mux.HandleFunc("/linecallback", handleLineRegister)
	// 		fmt.Println("Line ID is nill")
	// 	}
	// }
	return
}
