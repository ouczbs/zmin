package main

import (
	"github.com/ouczbs/Zmin/component/gate"
)

func main()  {
	service := gate.NewGateService()
	service.Run()
}