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
)

var (
	lineOauthConfig *oauth2.Config

	oauthStateStringLine = "random"

	lineID = ""
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	lineOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("LINE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("LINE_OAUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("LINE_OAUTH_CALLBACK"),
		Scopes:       []string{"profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://access.line.me/oauth2/v2.1/authorize?response_type=code",
			TokenURL: "https://api.line.me/oauth2/v2.1/token",
		},
	}
}

func handleLineLogin(w http.ResponseWriter, r *http.Request) {
	url := lineOauthConfig.AuthCodeURL(oauthStateStringLine)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	fmt.Println()
	fmt.Println("Complete")
	fmt.Println()
}

func handleLineCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserLine(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		fmt.Print(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var data UserLine
	json.Unmarshal(content, &data)
	// fmt.Printf("Line = %s\n", content)
	// fmt.Println()

	dbData := database.FetchData()

	fmt.Printf("Your User ID = %s\n", data.UserID)
	fmt.Println("=====================================================")
	fmt.Println()

	for _, el := range dbData {
		if data.UserID == *el.LineID {
			// fmt.Println("User ID =", data.UserID)
			// fmt.Println("DB_UserID =", el.LineID)
			// fmt.Println("true")
			// fmt.Println("--------------------------------------------")
			lineID = *el.LineID
			break
		} else {
			// fmt.Println("User ID =", data.UserID)
			// fmt.Println("DB_UserID =", el.LineID)
			// fmt.Println("false")
			// fmt.Println("--------------------------------------------")
			lineID = ""
		}
	}

	if lineID == "" {
		fmt.Println()
		fmt.Println("false")
		fmt.Println("User ID mismatch")
		fmt.Println("============================================")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}

	if lineID == data.UserID {
		fmt.Println()
		fmt.Println("true")
		fmt.Println("User ID is", lineID)
		fmt.Println("Your User ID is", data.UserID)
		fmt.Println("===========================================")
		http.ServeFile(w, r, "templates/success.html")
	}
}

func getUserLine(state string, code string) ([]byte, error) {
	if state != oauthStateStringLine {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := lineOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.line.me/v2/profile", nil)
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	response, _ := client.Do(req)

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}
