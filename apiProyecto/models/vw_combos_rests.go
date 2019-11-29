package models

type VistaCmb struct {
	Restaurant  string    `json:"restaurante,omitempty"`
	Combo 		 string    `json:"combo, omitempty"`
	Detail    string `json:"detail,omitempty"`
	ComboPrice	 int 	`json:"comboprice,omitempty"`
	IdRest int `json:"idrest,omitempty"`
}

