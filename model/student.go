package model

//Student - student dao
type Student struct {
	StudentID int    `json:"studentID,omitempty"`
	Name      string `json:"name,omitempty"`
	Password  string `json:"password,omitempty"`
}
