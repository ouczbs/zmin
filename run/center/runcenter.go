package main

import "github.com/ouczbs/Zmin/component/center"

func main()  {
	service := center.NewCenterService()
	service.Run()
}