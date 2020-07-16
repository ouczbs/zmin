package main

import (
	"Zmin/component/gate"
)

func main()  {
	service := gate.NewGateService()
	service.Run()
}