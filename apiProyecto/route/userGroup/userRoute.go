package userGroup

import (
	"apiProyecto/controller/userController"
	"github.com/labstack/echo"
)

func UserGroupAccess(g *echo.Group){
	g.GET("obtenerUsuarios", userController.VerUsers)
	g.POST("buscarUsuario", userController.SearchUser)
	g.POST("guardarUsuario", userController.InsertUsers)
}
