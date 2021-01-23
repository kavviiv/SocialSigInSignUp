package handlers

import (
	"Project/login/database"
	"net/http"
)

// New :
func New() http.Handler {
	mux := http.NewServeMux()
	// Root
	mux.Handle("/", http.FileServer(http.Dir("templates/")))

	data := database.CheckData()

	if data != nil {
		mux.HandleFunc("/google/login", handleGoogleLogin)
		mux.HandleFunc("/googlecallback", handleGoogleCallback)
		mux.HandleFunc("/facebook/login", handleFacebookLogin)
		mux.HandleFunc("/facebookcallback", handleFacebookCallback)
		mux.HandleFunc("/facebook/regist", handleFacebookRegister)
		mux.HandleFunc("/line/login", handleLineLogin)
		mux.HandleFunc("/linecallback", handleLineCallback)

	}

	if data == nil {
		mux.HandleFunc("/googlecallback", handleGoogleRegister)
	}

	// mux.HandleFunc("/line/regist", handleGoogleLogin)
	// mux.HandleFunc("/google/regist", handleGoogleLogin)

	return mux
}
