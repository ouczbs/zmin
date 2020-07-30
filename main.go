package main

import (
	"Zmin/component/center"
	"Zmin/engine/zcache"
	"Zmin/engine/zlog"
	"Zmin/engine/zmodel"
	"Zmin/engine/zproto/pb"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"os/exec"
	"strconv"
	"time"
)
var Service = zmodel.Service
type UService = zmodel.UService
func main() {
	zmodel.InitService()
	runCenter()
	runGate()
	runLogin()
	runDispatcher()
	for{
		time.Sleep(time.Duration(1))
	}
}
func runCenter(){
	var service UService
	err := zcache.MongoClient.FindOne(Service,bson.M{"type":pb.COMPONENT_TYPE_CENTER} , &service)
	if err!= nil{
		zlog.Debug(err)
		return
	}
	os.Args = []string{service.Path , "-ComponentId" , strconv.Itoa(int(service.Id)) , "-ListenAddr" , service.ListenAddr , service.Property}
	center := center.NewCenterService()
	go center.Run()
}
func runGate(){
	var service UService
	err := zcache.MongoClient.FindOne(Service,bson.M{"type":pb.COMPONENT_TYPE_GATE} , &service)
	if err!= nil{
		zlog.Debug(err)
		return
	}
	runService(&service)
}
func runLogin(){
	var serviceList []UService
	err := zcache.MongoClient.Find(Service,bson.M{"type":pb.COMPONENT_TYPE_LOGIN} , &serviceList)
	if err!= nil{
		zlog.Debug(err)
		return
	}
	for _,service := range serviceList{
		runService(&service)
	}
}
func runDispatcher(){
	var serviceList []UService
	err := zcache.MongoClient.Find(Service,bson.M{"type":pb.COMPONENT_TYPE_DISPATCHER} , &serviceList)
	if err!= nil{
		zlog.Debug(err)
		return
	}
	for _,service := range serviceList{
		runService(&service)
	}
}
func runService(service * UService)  {
	args := []string{"run",service.Path,"-ComponentId" , strconv.Itoa(int(service.Id)) , "-ListenAddr" , service.ListenAddr , service.Property}
	cmd := exec.Command("go", args...)
	err := cmd.Start()
	if err != nil {
		zlog.Errorf("run service failed" ,err)
	}
}