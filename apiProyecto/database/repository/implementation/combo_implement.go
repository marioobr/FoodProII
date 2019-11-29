package implementation

import (
	repo "apiProyecto/database/repository/interface"
	"apiProyecto/models"
	"database/sql"
	"errors"
)

type ComboRepoImpl struct {
	Db *sql.DB
}

func (c ComboRepoImpl) Select() ([]models.Combo, error) {
	combo_list := make([]models.Combo, 0)
	rows, err := c.Db.Query("Select id_combo, id_rest,combo_name, combo_description, combo_price from combo")
	if err != nil {
		return combo_list, err
	}
	for rows.Next() {
		combo := models.Combo{}
		err := rows.Scan(&combo.IdCombo, &combo.IdRest, &combo.ComboName, &combo.ComboDesc, &combo.ComboPrice)
		if err != nil {
			break
		}
		combo_list = append(combo_list, combo)
	}
	err = rows.Err()
	if err != nil {
		return combo_list, err
	}
	return combo_list, nil
}

func (c ComboRepoImpl) Insert(combo *models.Combo) (bool, error) {
	// new, err := u.Db.Query("Insert into users (username,password) values ($1,$2)", user.Username, user.Password)

	// if err != nil {
	// 	return false, err
	// 	new.Close()
	// }

	query, err := c.Db.Prepare("Insert into combo (combo_name, combo_description, combo_price, id_rest, state) values ($1,$2,$3,$4,1)")
	if err != nil {
		return false, err
	}

	result, err := query.Exec(combo.ComboName, combo.ComboDesc, combo.ComboPrice, combo.IdRest)
	if err != nil {
		return false, err
	}

	i, err := result.RowsAffected()
	if i != 1 {
		return false, errors.New("Error: Muchas filas afectadas")
	}

	return true, nil
}

func (c ComboRepoImpl) Filter(combo *models.Combo) ([]models.Combo, bool, error) {
	combo_list := make([]models.Combo, 0)
	rows, err := c.Db.Query("Select id_rest, combo_name, combo_description, combo_price from combo where combo_name = $1", combo.ComboName)
	if err != nil {
		return combo_list, false, err
	}
	for rows.Next() {
		combos := models.Combo{}
		err := rows.Scan(&combos.IdRest, &combos.ComboName, &combos.ComboDesc, &combos.ComboPrice)
		if err != nil {
			break
		}
		combo_list = append(combo_list, combos)
	}
	err = rows.Err()
	if err != nil {
		return combo_list, false, err
	}
	return combo_list, true, nil
}

func NewComboRepoImpl(db *sql.DB) repo.ComboInterface {
	return &ComboRepoImpl{Db: db}
}
