package main

import "github.com/ouczbs/zmin/component/center"

func main()  {
	service := center.NewCenterService()
	service.Run()
}