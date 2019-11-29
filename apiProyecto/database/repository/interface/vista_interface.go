package _interface

import "apiProyecto/models"

type VistaInterface interface {
	Filter(combo *models.VistaCmb) ([]models.VistaCmb, bool, error)
}
