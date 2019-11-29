package models


type User struct {
	IdUser int `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Names string `json:"names,omitempty"`
	LastNames string `json:"lastnames,omitempty"`
}
