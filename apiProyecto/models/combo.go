package models

type Combo struct {
	IdCombo      int    `json:"idcombo,omitempty"`
	IdRest 		 int    `json:"idrest, omitempty"`
	ComboName    string `json:"comboname,omitempty"`
	ComboPrice	 int 	`json:"comboprice,omitempty"`
	ComboDesc    string `json:"combodesc,omitempty"`
}
