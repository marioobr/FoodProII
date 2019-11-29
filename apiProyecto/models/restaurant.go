package models

type Restaurant struct {
	IdRest      int    `json:"idrest,omitempty"`
	RestName    string `json:"restname,omitempty"`
	RestAddress string `json:"restaddress,omitempty"`
	RestPhone   int    `json:"restphone,omitempty"`
	RestImage   string `json:"restimage,omitempty"`
}
