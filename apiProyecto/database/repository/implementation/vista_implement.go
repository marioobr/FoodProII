package implementation

import (
	repo "apiProyecto/database/repository/interface"
	"apiProyecto/models"
	"database/sql"
)
type VistaRepoImpl struct {
	Db *sql.DB
}

func (c VistaRepoImpl) Filter(combo *models.VistaCmb) ([]models.VistaCmb, bool, error) {
	combo_list := make([]models.VistaCmb, 0)
	rows, err := c.Db.Query("Select restaurante, combo, detalle, precio from vw_combos_rests where id_restaurante = $1", combo.IdRest)
	if err != nil {
		return combo_list, false, err
	}
	for rows.Next() {
		combo := models.VistaCmb{}
		err := rows.Scan(&combo.Restaurant, &combo.Combo, &combo.Detail, &combo.ComboPrice)
		if err != nil {
			break
		}
		combo_list = append(combo_list, combo)
	}
	err = rows.Err()
	if err != nil {
		return combo_list, false, err
	}
	return combo_list, true, nil
}

func NewVistaRepoImpl(db *sql.DB) repo.VistaInterface {
	return &VistaRepoImpl{Db: db}
}