package database

// UserData :
type UserData struct {
	GoogleID   string `json:"googleID"`
	FacebookID string `json:"facebookID"`
	LineID     string `json:"lineID"`
	Email      string `json:"email"`
}
