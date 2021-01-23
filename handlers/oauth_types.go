package handlers

// UserGoogle :
type UserGoogle struct {
	UserID string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

// UserFacebook :
type UserFacebook struct {
	UserID string `json:"id"`
	// Name   string `json:"name"`
	Email string `json:"email"`
}

// UserLine :
type UserLine struct {
	UserID string `json:"userId"`
	Name   string `json:"displayName"`
}
