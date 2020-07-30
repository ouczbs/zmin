package zmodel

import (
	"github.com/ouczbs/Zmin/engine/zattr"
	"github.com/ouczbs/Zmin/engine/zcache"
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zlog"
	"github.com/ouczbs/Zmin/engine/zproto/pb"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"strconv"
	"strings"
)

func InitService(){
	zcache.MongoClient.ClearTable(Service)
	initCenterService()
	initLoginService()
	initGateService()
	initDispatcherService()
}
var sequence zconf.TSequence = 0
func Sequence()zconf.TSequence{
	sequence++
	return sequence
}
func writeString(bytes *strings.Builder , s string){
	bytes.WriteString(s)
	bytes.WriteString(" ")
}
func writeCenterAddrProperty(bytes *strings.Builder , listenAddr string){
	writeString(bytes , strconv.Itoa(int(zattr.StringCenterAddr)))
	writeString(bytes , strconv.Itoa(int(pb.Property_Type_String)))
	writeString(bytes , listenAddr)
}
func writeLoginProperty(bytes *strings.Builder , name string , centerAddr string)string{
	writeString(bytes , strconv.Itoa(int(zattr.StringComponentName)))
	writeString(bytes , strconv.Itoa(int(pb.Property_Type_String)))
	writeString(bytes , name)
	writeCenterAddrProperty(bytes , centerAddr)
	return bytes.String()
}
func writeBaseProperty(bytes *strings.Builder , centerAddr string)string{
	writeCenterAddrProperty(bytes , centerAddr)
	return bytes.String()
}

func initLoginService(){
	//config := zconf.LoginConfig
	centerAddr := zconf.CenterConfig.ListenAddr
	dir,_ := os.Getwd()
	dir = strings.Replace(dir,"\\engine\\zmodel","",1)
	path := dir + "\\run\\login\\runlogin.go"
	//127.0.0.1:11001
	listenAddrList := [3]string{"127.0.0.1:11001", "127.0.0.1:11002", "127.0.0.1:11003"}
	nameList := [3]string{"一区", "二区", "三区"}
	for i,listenAddr := range listenAddrList{
		name := nameList[i]
		var bytes strings.Builder
		service := &UService{
			Id: Sequence(),
			ListenAddr: listenAddr,
			Type: pb.COMPONENT_TYPE_LOGIN,
			Path:path,
			Property: writeLoginProperty(&bytes , name , centerAddr),
		}
		zcache.MongoClient.UpdateOrInsert(service,bson.M{"id":service.Id})
	}
}
func initCenterService(){
	centerAddr := zconf.CenterConfig.ListenAddr
	dir,_ := os.Getwd()
	dir = strings.Replace(dir,"\\engine\\zmodel","",1)
	path := dir + "\\run\\center\\runcenter.go"
	service := &UService{
		Id: Sequence(),
		ListenAddr: centerAddr,
		Type: pb.COMPONENT_TYPE_CENTER,
		Path:path,
	}
	zcache.MongoClient.UpdateOrInsert(service,bson.M{"id":service.Id})
}
func initGateService(){
	centerAddr := zconf.CenterConfig.ListenAddr
	gateAddr := zconf.GateConfig.ListenAddr
	dir,_ := os.Getwd()
	dir = strings.Replace(dir,"\\engine\\zmodel","",1)
	path := dir + "\\run\\gate\\rungate.go"
	var bytes strings.Builder
	service := &UService{
		Id: Sequence(),
		ListenAddr: gateAddr,
		Type: pb.COMPONENT_TYPE_GATE,
		Path:path,
		Property: writeBaseProperty(&bytes , centerAddr),
	}
	zcache.MongoClient.UpdateOrInsert(service,bson.M{"id":service.Id})
}
func initDispatcherService(){
	centerAddr := zconf.CenterConfig.ListenAddr
	dir,_ := os.Getwd()
	dir = strings.Replace(dir,"\\engine\\zmodel","",1)
	path := dir + "\\run\\dispatcher\\rundispatcher.go"
	listenAddrList := [3]string{"127.0.0.1:12001", "127.0.0.1:12002", "127.0.0.1:12003"}
	for _,listenAddr := range listenAddrList{
		var bytes strings.Builder
		service := &UService{
			Id: Sequence(),
			ListenAddr: listenAddr,
			Type: pb.COMPONENT_TYPE_DISPATCHER,
			Path:path,
			Property: writeBaseProperty(&bytes , centerAddr),
		}
		zcache.MongoClient.UpdateOrInsert(service,bson.M{"id":service.Id})
	}
}
func readLoginService(){
	var results []UService
	zcache.MongoClient.Find(Service,bson.M{"type":pb.COMPONENT_TYPE_LOGIN}, &results)
	zlog.Debug(results)
}