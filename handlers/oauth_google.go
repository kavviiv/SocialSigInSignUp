package handlers

import (
	"Project/login/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config

	oauthStateStringGoogle = "random"

	googleEmail = ""
)

func config() {
	err := godotenv.Load()
	log.Printf("%s", err)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// เพิ่ม redirecturl ของ register
	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_OAUTH_CALLBACK"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"openid"},
		Endpoint: google.Endpoint,
	}
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	config()
	url := googleOauthConfig.AuthCodeURL(oauthStateStringGoogle)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	fmt.Println()
	fmt.Println("Google login success")
	fmt.Println()
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	config()
	fmt.Println("Google Login Callback")
	content, err := getUserGoogle(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var data = UserGoogle{}
	json.Unmarshal(content, &data)

	dbData := database.FetchData()

	fmt.Println()
	fmt.Printf("Your User ID = %s\n", data.UserID)
	fmt.Printf("Your Email = %s\n", data.Email)
	fmt.Println("=====================================================")
	fmt.Println()

	for _, el := range dbData {
		if el.Email != nil {
			if data.Email == *el.Email {
				http.ServeFile(w, r, "templates/mainPage.html")
				break
			} else {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			}
		}
	}
}

func handleGoogleRegister(w http.ResponseWriter, r *http.Request) {
	config()
	url := googleOauthConfig.AuthCodeURL(oauthStateStringGoogle)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	fmt.Println()
	fmt.Println("Google register")
	fmt.Println()
}

func handleGoogleRegisterCallback(w http.ResponseWriter, r *http.Request) {
	config()
	fmt.Println("Google register callback")
	content, err := getUserGoogle(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var data = UserGoogle{}
	json.Unmarshal(content, &data)
	fmt.Printf("Register Data = %s\n", content)

	db := database.OpenConn()

	sqlStatement := `UPDATE test SET google_id = $1, email = $2 WHERE user_id='654321'`
	_, err = db.Exec(sqlStatement, data.UserID, data.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	fmt.Println("Google register success")
	http.ServeFile(w, r, "templates/mainPage.html")
	w.WriteHeader(http.StatusOK)
	defer db.Close()
	return
}

func getUserGoogle(state string, code string) ([]byte, error) {
	config()
	if state != oauthStateStringGoogle {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
