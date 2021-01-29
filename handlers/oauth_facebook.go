package handlers

import (
	"Project/login/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var (
	facebookOauthConfig *oauth2.Config

	oauthStateStringFacebook = "random"

	facebookID = ""
)

func init() {
	err := godotenv.Load()
	log.Printf("%s", err)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	facebookOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("FACEBOOK_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("FACEBOOK_OAUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("FACEBOOK_OAUTH_CALLBACK"),
		Scopes:       []string{"public_profile", "email"},
		Endpoint:     facebook.Endpoint,
	}
}

func handleFacebookLogin(w http.ResponseWriter, r *http.Request) {
	url := facebookOauthConfig.AuthCodeURL(oauthStateStringFacebook)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	fmt.Println()
	fmt.Println("Complete")
	fmt.Println()
}

func handleFacebookCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateStringFacebook {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateStringFacebook, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")

	token, err := facebookOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://graph.facebook.com/me?fields=email&access_token=" +
		url.QueryEscape(token.AccessToken))
	if err != nil {
		fmt.Printf("Get: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var data UserFacebook
	json.Unmarshal(response, &data)
	// fmt.Printf("Facebook = %s\n", response)

	dbData := database.FetchData()

	fmt.Println()
	fmt.Printf("Your User ID = %s\n", data.UserID)
	fmt.Println("=====================================================")
	fmt.Println()

	for _, el := range dbData {
		if el.FacebookID != nil {
			if data.UserID == *el.FacebookID {
				// fmt.Println("User ID =", data.UserID)
				// fmt.Println("DB_UserID =", el.FacebookID)
				// fmt.Println("true")
				// fmt.Println("--------------------------------------------")
				http.ServeFile(w, r, "templates/mainPage.html")
				break
			} else {
				// fmt.Println("User ID =", data.UserID)
				// fmt.Println("DB_UserID =", el.FacebookID)
				// fmt.Println("false")
				// fmt.Println("--------------------------------------------")
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			}
		}
	}
}

func handleFacebookRegister(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateStringFacebook {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateStringFacebook, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")

	token, err := facebookOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://graph.facebook.com/me?fields=email&access_token=" +
		url.QueryEscape(token.AccessToken))
	if err != nil {
		fmt.Printf("Get: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var data UserFacebook
	json.Unmarshal(response, &data)
	// // fmt.Printf("Facebook = %s\n", response)

	db := database.OpenConn()

	// var data UserFacebook
	// err := json.NewDecoder(r.Body).Decode(&data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	sqlQuery := `INSERT INTO test (facebook_id) VALUES ($1) WHERE user_id='654321'`
	_, err = db.Exec(sqlQuery, data.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}
