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

	facebookID                       = ""
	facebook_signin, facebook_regist string
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
		RedirectURL:  os.Getenv("oauthURI") + "facebookcallback",
		Scopes:       []string{"public_profile", "email"},
		Endpoint:     facebook.Endpoint,
	}
}

func handleFacebookLogin(w http.ResponseWriter, r *http.Request) {
	url := facebookOauthConfig.AuthCodeURL(oauthStateStringFacebook)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	fmt.Println()
	fmt.Println("Facebook login success")
	fmt.Println()
	facebook_signin = "true"
}

func handleFacebookRegister(w http.ResponseWriter, r *http.Request) {
	url := facebookOauthConfig.AuthCodeURL(oauthStateStringFacebook)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	fmt.Println()
	fmt.Println("Facebook regist")
	fmt.Println()
	facebook_regist = "true"
}

func handleFacebookCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Facebook login callback")
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
	if facebook_signin == "true" {
		var data UserFacebook
		json.Unmarshal(response, &data)

		dbData := database.FetchData()

		fmt.Println()
		fmt.Printf("Your User ID = %s\n", data.UserID)
		fmt.Println("=====================================================")
		fmt.Println()

		for _, el := range dbData {
			if el.FacebookID != nil {
				if data.UserID == *el.FacebookID {
					http.ServeFile(w, r, "templates/mainPage.html")
					break
				} else {
					http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				}
			}
		}

	}

	if facebook_regist == "true" {
		var data UserFacebook
		json.Unmarshal(response, &data)

		db := database.OpenConn()

		sqlQuery := `UPDATE test SET facebook_id = $1 WHERE user_id='123456'`
		_, err = db.Exec(sqlQuery, data.UserID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			panic(err)
		}

		fmt.Println("Facebook register success")
		http.Redirect(w, r, "/mainPage.html", 307)
		//http.ServeFile(w, r, "templates/mainPage.html")
		//w.WriteHeader(http.StatusOK)
		defer db.Close()

	}

}

// func handleFacebookRegister(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Facebook register callback")
// 	state := r.FormValue("state")
// 	if state != oauthStateStringFacebook {
// 		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateStringFacebook, state)
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}

// 	code := r.FormValue("code")

// 	token, err := facebookOauthConfig.Exchange(oauth2.NoContext, code)
// 	if err != nil {
// 		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}

// 	resp, err := http.Get("https://graph.facebook.com/me?fields=email&access_token=" +
// 		url.QueryEscape(token.AccessToken))
// 	if err != nil {
// 		fmt.Printf("Get: %s\n", err)
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	response, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Printf("ReadAll: %s\n", err)
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}

// 	var data UserFacebook
// 	json.Unmarshal(response, &data)

// 	db := database.OpenConn()

// 	sqlQuery := `INSERT INTO test (facebook_id) VALUES ($1) WHERE user_id='654321'`
// 	_, err = db.Exec(sqlQuery, data.UserID)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		panic(err)
// 	}

// 	fmt.Println("Facebook register success")
// 	http.ServeFile(w, r, "templates/mainPage.html")
// 	w.WriteHeader(http.StatusOK)
// 	defer db.Close()
// }
