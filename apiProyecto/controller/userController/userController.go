package userController

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

func VerUsers(c echo.Context) error {
	userRepository := dbImplement.NewUserRepoImpl(driver.Instance().SQL)
	users, err := userRepository.Select()
	if err != nil {
		fmt.Println(err)
	}
	return response.Success(c, users)
}

func InsertUsers(c echo.Context) error {
	// user := models.User{
	// 	Username: c.FormValue("username"),
	// }
	user := new(models.User)
	e := c.Bind(user)
	if e != nil {
		return response.Error(c, http.StatusBadRequest, "Error en Json")
	}
	userRepository := dbImplement.NewUserRepoImpl(driver.Instance().SQL)
	users, err := userRepository.Insert(user)
	if err != nil {
		fmt.Println(err)
	}
	if users {
		return response.SuccessWithCode(c, 200, user)
	}
	return errors.New("No se ingresó el usuario")
}

func SearchUser(c echo.Context) error {
	// user := models.User{
	// 	Username: c.FormValue("username"),
	// }
	user := new(models.User)
	e := c.Bind(user)
	if e != nil {
		return response.Error(c, http.StatusBadRequest, "Error en Json")
	}
	userRepository := dbImplement.NewUserRepoImpl(driver.Instance().SQL)
	users, found, err := userRepository.Filter(user)
	if err != nil {
		fmt.Println(err)
	}
	if !found {
		return response.Error(c, http.StatusBadRequest, "Usuario no existe")
	}
	return response.Success(c, users)
}