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

	lineID                   = ""
	line_signin, line_regist string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	lineOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("LINE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("LINE_OAUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("oauthURI") + "linecallback",
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
	fmt.Println("Line login success")
	fmt.Println()
	line_signin = "true"
}

func handleLineRegister(w http.ResponseWriter, r *http.Request) {
	url := lineOauthConfig.AuthCodeURL(oauthStateStringLine)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	fmt.Println("Regist Line")
	line_regist = "true"
}

func handleLineCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Line callback")
	fmt.Println(line_regist)
	content, err := getUserLine(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		fmt.Print(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		fmt.Println("l")
		return
	}
	if line_signin == "true" {
		var data UserLine
		json.Unmarshal(content, &data)
		dbData := database.FetchData()
		fmt.Printf("Your User ID = %s\n", data.UserID)
		fmt.Println("=====================================================")
		fmt.Println()

		for _, el := range dbData {
			if el.LineID != nil {
				if data.UserID == *el.LineID {
					http.ServeFile(w, r, "templates/mainPage.html")
					break
				} else {
					http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				}
			}
		}
	}

	if line_regist == "true" {
		var data = UserLine{}
		json.Unmarshal(content, &data)
		db := database.OpenConn()
		sqlStatement := `UPDATE test SET line_id = $1 WHERE user_id='123456'`
		_, err = db.Exec(sqlStatement, data.UserID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			panic(err)
		}
		fmt.Println("Line register success")
		http.ServeFile(w, r, "templates/mainPage.html")
		w.WriteHeader(http.StatusOK)
		defer db.Close()
		return
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
