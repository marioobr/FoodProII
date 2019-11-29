package _interface

import "apiProyecto/models"

type ComboInterface interface {
	Select() ([]models.Combo, error)
	Filter(combo *models.Combo) ([]models.Combo, bool, error)
	Insert(combo *models.Combo) (bool, error)
}