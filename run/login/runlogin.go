package main

import (
	"github.com/ouczbs/Zmin/component/login"
)

func main()  {
	service := login.NewLoginService()
	service.Run()
}