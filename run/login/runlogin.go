package main

import (
	"github.com/ouczbs/zmin/component/login"
)

func main()  {
	service := login.NewLoginService()
	service.Run()
}