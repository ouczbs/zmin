package main

import (
	"github.com/ouczbs/zmin/component/dispatcher"
)

func main()  {
	service := dispatcher.NewDispatcherService()
	service.Run()
}