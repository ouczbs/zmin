package main

import (
	"Zmin/component/dispatcher"
)

func main()  {
	service := dispatcher.NewDispatcherService()
	service.Run()
}