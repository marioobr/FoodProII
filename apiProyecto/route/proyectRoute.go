package route

import (
	"github.com/labstack/echo"
	rg "apiProyecto/route/restGroup"
	ug "apiProyecto/route/userGroup"
)
func Init() *echo.Echo {
	e := echo.New()
	restGroup := e.Group("rest/")
	userGroup := e.Group("user/")

	rg.RestGroupAccess(restGroup)
	ug.UserGroupAccess(userGroup)

	return e
}