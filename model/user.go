package model

//User - student dao
type User struct {
	UserID   int    `json:"userID,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}
