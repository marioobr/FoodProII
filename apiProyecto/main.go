package main

import (
	"apiProyecto/route"
	"apiProyecto/utils"
)

func main() {
	e := route.Init()
	port := utils.GetPort()
	e.Logger.Fatal(e.Start(port))
}