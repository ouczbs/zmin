package main

import "Zmin/component/center"

func main()  {
	service := center.NewCenterService()
	service.Run()
}