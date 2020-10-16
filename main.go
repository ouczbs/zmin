package main

import (
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/ouczbs/Zmin/component/center"
	"github.com/ouczbs/Zmin/engine/zcache"
	"github.com/ouczbs/Zmin/engine/zlog"
	"github.com/ouczbs/Zmin/engine/zmodel"
	"github.com/ouczbs/Zmin/engine/zproto/zpb"
	"go.mongodb.org/mongo-driver/bson"
)

var Service = zmodel.Service

type UService = zmodel.UService

func main() {
	zcache.InitMongoClient("mongodb://111.229.54.9:27017", "mmo")
	zmodel.InitService()
	runCenter()
	//runGate()
	runLogin()
	runDispatcher()
	for {
		time.Sleep(time.Duration(1))
	}
}
func runCenter() {
	var service UService
	err := zcache.GetMongoClient().FindOne(&service, bson.M{"type": zpb.COMPONENT_TYPE_CENTER})
	if err != nil {
		zlog.Debug(err)
		return
	}
	os.Args = []string{service.Path, "-ComponentId", strconv.Itoa(int(service.Id)), "-ListenAddr", service.ListenAddr, service.Property}
	center := center.NewCenterService()
	go center.Run()
}
func runGate() {
	var service UService
	err := zcache.GetMongoClient().FindOne(&service, bson.M{"type": zpb.COMPONENT_TYPE_GATE})
	if err != nil {
		zlog.Debug(err)
		return
	}
	runService(&service)
}
func runLogin() {
	var serviceList []UService
	err := zcache.GetMongoClient().Find(Service, bson.M{"type": zpb.COMPONENT_TYPE_LOGIN}, &serviceList)
	if err != nil {
		zlog.Debug(err)
		return
	}
	for _, service := range serviceList {
		runService(&service)
	}
}
func runDispatcher() {
	var serviceList []UService
	err := zcache.GetMongoClient().Find(Service, bson.M{"type": zpb.COMPONENT_TYPE_DISPATCHER}, &serviceList)
	if err != nil {
		zlog.Debug(err)
		return
	}
	for _, service := range serviceList {
		runService(&service)
	}
}
func runService(service *UService) {
	args := []string{"run", service.Path, "-ComponentId", strconv.Itoa(int(service.Id)), "-ListenAddr", service.ListenAddr, service.Property}
	cmd := exec.Command("go", args...)
	err := cmd.Start()
	if err != nil {
		zlog.Errorf("run service failed", err)
	}
}
