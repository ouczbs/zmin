package main

import (
	"github.com/ouczbs/zmin/component/center"
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/data/zcache"
	"github.com/ouczbs/zmin/engine/data/zmodel"
	"github.com/ouczbs/zmin/engine/sync/zpb"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var Service = zmodel.Service

type FService = zmodel.FService

func main() {
	zcache.InitMongoClient("mongodb://124.221.147.27:27017", "mmo")
	InitService()
	runCenter()
	//runGate()
	runLogin()
	runDispatcher()
	for {
		time.Sleep(time.Duration(1))
	}
}
func runCenter() {
	var service FService
	err := zcache.GetMongoClient().FindOne(&service, bson.M{"type": zpb.COMPONENT_TYPE_CENTER})
	if err != nil {
		zlog.Error(err)
		return
	}
	os.Args = []string{service.Path, "-ComponentId", strconv.Itoa(int(service.Id)), "-ListenAddr", service.ListenAddr, service.Property}
	center := center.NewCenterService()
	go center.Run()
}
func runGate() {
	var service FService
	err := zcache.GetMongoClient().FindOne(&service, bson.M{"type": zpb.COMPONENT_TYPE_GATE})
	if err != nil {
		zlog.Error(err)
		return
	}
	runService(&service)
}
func runLogin() {
	var serviceList []FService
	err := zcache.GetMongoClient().Find(Service, bson.M{"type": zpb.COMPONENT_TYPE_LOGIN}, &serviceList)
	if err != nil {
		zlog.Error(err)
		return
	}
	for _, service := range serviceList {
		runService(&service)
	}
}
func runDispatcher() {
	var serviceList []FService
	err := zcache.GetMongoClient().Find(Service, bson.M{"type": zpb.COMPONENT_TYPE_DISPATCHER}, &serviceList)
	if err != nil {
		zlog.Error(err)
		return
	}
	for _, service := range serviceList {
		runService(&service)
	}
}
func runService(service *FService) {
	args := []string{"run", service.Path, "-ComponentId", strconv.Itoa(int(service.Id)), "-ListenAddr", service.ListenAddr, service.Property}
	cmd := exec.Command("go", args...)
	err := cmd.Start()
	if err != nil {
		zlog.Errorf("run service failed", err)
	}
}
