package main

import (
	"github.com/ouczbs/Zmin/component/dispatcher"
)

func main()  {
	service := dispatcher.NewDispatcherService()
	service.Run()
}