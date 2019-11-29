package _interface

import "apiProyecto/models"

type UserInterface interface {
	Select() ([]models.User, error)
	Filter(user *models.User) ([]models.User, bool, error)
	Insert(user *models.User) (bool, error)
}