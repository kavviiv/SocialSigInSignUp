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

func init() {
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
	url := googleOauthConfig.AuthCodeURL(oauthStateStringGoogle)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	fmt.Println()
	fmt.Println("Login")
	fmt.Println()
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Google Login Callback")
	content, err := getUserGoogle(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var data = UserGoogle{}
	json.Unmarshal(content, &data)
	// fmt.Printf("Google = %s\n", content)
	// fmt.Println(reflect.ValueOf(data).Kind())

	// fmt.Printf("Data = %s\n", content)
	// fmt.Println()

	dbData := database.FetchData()

	fmt.Println()
	fmt.Printf("Your User ID = %s\n", data.UserID)
	fmt.Printf("Your Email = %s\n", data.Email)
	fmt.Println("=====================================================")
	fmt.Println()

	for _, el := range dbData {
		if el.Email != nil {
			if data.Email == *el.Email {
				// fmt.Println("UserID =", data.Email)
				// fmt.Println("DB_Email =", el.Email)
				// fmt.Println("true")
				// fmt.Println("--------------------------------------------")
				// googleEmail = *el.Email
				http.ServeFile(w, r, "templates/ConnectSocial.html")
				break
			} else {
				// fmt.Println("Eamil =", data.Email)
				// fmt.Println("DB_Email =", el.Email)
				// fmt.Println("false")
				// fmt.Println("--------------------------------------------")
				// googleEmail = ""
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			}
		}
	}

	// if googleEmail == "" {
	// 	fmt.Println()
	// 	fmt.Println("false")
	// 	fmt.Println("Your Email mismatch")
	// 	fmt.Println("============================================")
	// 	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// }

	// if googleEmail == data.Email {
	// 	fmt.Println()
	// 	fmt.Println("true")
	// 	fmt.Println("Email is", googleEmail)
	// 	fmt.Println("User Eamil is", data.Email)
	// 	fmt.Println("===========================================")
	// 	http.ServeFile(w, r, "templates/ConnectSocial.html")
	// }
}

func handleGoogleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Google Register")
	content, err := getUserGoogle(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var data = UserGoogle{}
	json.Unmarshal(content, &data)

	db := database.OpenConn()

	sqlStatement := `INSERT INTO person (google_id) VALUES ($1) WHERE user_id='654321'`
	_, err = db.Exec(sqlStatement, data.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
	return
}

func getUserGoogle(state string, code string) ([]byte, error) {

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
