package restGroup

import (
	"apiProyecto/controller"
	"github.com/labstack/echo"
)

func RestGroupAccess(g *echo.Group) {
	g.GET("obtenerRests", controller.VerRests)
	g.POST("guardarRest", controller.InsertRests)
	g.GET("obtenerCombos", controller.VerCombos)
	g.POST("guardarCombos", controller.InsertCombos)
	g.POST("vistaCombosRes", controller.SpecificCombo)
}
