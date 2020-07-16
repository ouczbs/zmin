package main

import (
	"Zmin/component/login"
)

func main()  {
	service := login.NewLoginService()
	service.Run()
}