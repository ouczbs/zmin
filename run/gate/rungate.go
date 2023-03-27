package main

import (
	"github.com/ouczbs/zmin/component/gate"
	"os"
	"strconv"
)

func main()  {
	path,_ := os.Getwd()
	os.Args = []string{path, "-ComponentId", strconv.Itoa(int(333)), "-ListenAddr", "127.0.0.1:13001", ""}
	service := gate.NewGateService()
	service.Run()
}