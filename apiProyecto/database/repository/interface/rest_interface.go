package _interface

import "apiProyecto/models"

type RestInterface interface {
	Select() ([]models.Restaurant, error)
	Filter(restaurant *models.Restaurant) ([]models.Restaurant, bool, error)
	Insert(restaurant *models.Restaurant) (bool, error)
}
