package controller

import (
	"apiProyecto/components/response"
	"apiProyecto/database/driver"
	dbImplement "apiProyecto/database/repository/implementation"
	"apiProyecto/models"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func VerRests(c echo.Context) error {
	restRepository := dbImplement.NewRestRepoImpl(driver.Instance().SQL)
	rests, err := restRepository.Select()
	if err != nil {
		fmt.Println(err)
	}
	return response.Success(c, rests)
}

func InsertRests(c echo.Context) error {
	// user := models.User{
	// 	Username: c.FormValue("username"),
	// }
	rest := new(models.Restaurant)
	e := c.Bind(rest)
	if e != nil {
		return response.Error(c, http.StatusBadRequest, "Error en Json")
	}
	restRepository := dbImplement.NewRestRepoImpl(driver.Instance().SQL)
	rests, err := restRepository.Insert(rest)
	if err != nil {
		fmt.Println(err)
	}
	if rests {
		return response.SuccessWithCode(c, 200, rest)
	}
	return errors.New("No se ingresó el restaurante")
}

func SearchRest(c echo.Context) error {
	// user := models.User{
	// 	Username: c.FormValue("username"),
	// }
	rest := new(models.Restaurant)
	e := c.Bind(rest)
	if e != nil {
		return response.Error(c, http.StatusBadRequest, "Error en Json")
	}
	restRepository := dbImplement.NewRestRepoImpl(driver.Instance().SQL)
	rests, found, err := restRepository.Filter(rest)
	if err != nil {
		fmt.Println(err)
	}
	if !found {
		return response.Error(c, http.StatusBadRequest, "Restaurante no existe")
	}
	return response.Success(c, rests)
}

func VerCombos(c echo.Context) error {
	comboRepository := dbImplement.NewComboRepoImpl(driver.Instance().SQL)
	combos, err := comboRepository.Select()
	if err != nil {
		fmt.Println(err)
	}
	return response.Success(c, combos)
}

func InsertCombos(c echo.Context) error {
	// user := models.User{
	// 	Username: c.FormValue("username"),
	// }
	combo := new(models.Combo)
	e := c.Bind(combo)
	if e != nil {
		return response.Error(c, http.StatusBadRequest, "Error en Json")
	}
	comboRepository := dbImplement.NewComboRepoImpl(driver.Instance().SQL)
	combos, err := comboRepository.Insert(combo)
	if err != nil {
		fmt.Println(err)
	}
	if combos {
		return response.SuccessWithCode(c, 200, combo)
	}
	return errors.New("No se ingresó el restaurante")
}

func SpecificCombo(c echo.Context) error {
	// user := models.User{
	// 	Username: c.FormValue("username"),
	// }
	combo := new(models.VistaCmb)
	e := c.Bind(combo)
	if e != nil {
		return response.Error(c, http.StatusBadRequest, "Error en Json")
	}
	cmbRepository := dbImplement.NewVistaRepoImpl(driver.Instance().SQL)
	cmbs, found, err := cmbRepository.Filter(combo)
	if err != nil {
		fmt.Println(err)
	}
	if !found {
		return response.Error(c, http.StatusBadRequest, "Restaurante no existe")
	}
	return response.Success(c, cmbs)
}