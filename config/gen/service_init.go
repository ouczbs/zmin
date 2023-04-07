package main

import (
	"encoding/json"
	"fmt"
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/core/zutil"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"os"
	"path/filepath"
	"strconv"
)

type (
	FServiceConfigFile = base.FServiceConfigFile
	FServiceConfig     = base.FServiceConfig
	FKVDBConfig        = base.FKVDBConfig
)

var (
	versionPort    = 9000
	loginPort      = 10000
	centerPort     = 11000
	gatePort       = 12000
	dispatcherPort = 13000
)

func toAddr(ip string, port int) string {
	return ip + ":" + strconv.Itoa(port)
}
func initDB() *FKVDBConfig {
	return &FKVDBConfig{
		Url:    "mongodb://124.221.147.27:27017",
		DB:     "mmo",
		Driver: "mongodb",
	}
}
func initVersion(ip string) *FServiceConfig {
	return &FServiceConfig{
		ComponentType: zconf.COMPONENT_TYPE_VERSION,
		ComponentId:   zutil.IncSequence(),
		ListenAddr:    toAddr(ip, versionPort),
		LogFile:       "version.log",
		LogLevel:      "debug",
	}
}
func initLoginList(ip string, version *FServiceConfig) []FServiceConfig {
	return []FServiceConfig{
		{
			ComponentType: zconf.COMPONENT_TYPE_LOGIN,
			ComponentId:   zutil.IncSequence(),
			ListenAddr:    toAddr(ip, loginPort),
			OwnerAddr:     version.ListenAddr,
			LogFile:       "login1.log",
			LogLevel:      "debug",
		},
		{
			ComponentType: zconf.COMPONENT_TYPE_LOGIN,
			ComponentId:   zutil.IncSequence(),
			ListenAddr:    toAddr(ip, loginPort+1),
			OwnerAddr:     version.ListenAddr,
			LogFile:       "login2.log",
			LogLevel:      "debug",
		},
	}
}
func initCenterList(ip string, version *FServiceConfig) []FServiceConfig {
	return []FServiceConfig{
		{
			ComponentType: zconf.COMPONENT_TYPE_CENTER,
			ComponentId:   zutil.IncSequence(),
			ListenAddr:    toAddr(ip, centerPort),
			OwnerAddr:     version.ListenAddr,
			LogFile:       "center1.log",
			LogLevel:      "debug",
		},
	}
}
func initGateList(ip string, center *FServiceConfig) []FServiceConfig {
	return []FServiceConfig{
		{
			ComponentType: zconf.COMPONENT_TYPE_GATE,
			ComponentId:   zutil.IncSequence(),
			ListenAddr:    toAddr(ip, gatePort),
			OwnerAddr:     center.ListenAddr,
			LogFile:       "gate1.log",
			LogLevel:      "debug",
		},
		{
			ComponentType: zconf.COMPONENT_TYPE_GATE,
			ComponentId:   zutil.IncSequence(),
			ListenAddr:    toAddr(ip, gatePort+1),
			OwnerAddr:     center.ListenAddr,
			LogFile:       "gate2.log",
			LogLevel:      "debug",
		},
	}
}
func initDispatcherList(ip string, center *FServiceConfig) []FServiceConfig {
	return []FServiceConfig{
		{
			ComponentType: zconf.COMPONENT_TYPE_DISPATCHER,
			ComponentId:   zutil.IncSequence(),
			ListenAddr:    toAddr(ip, dispatcherPort),
			OwnerAddr:     center.ListenAddr,
			LogFile:       "dispatcher1.log",
			LogLevel:      "debug",
		},
		{
			ComponentType: zconf.COMPONENT_TYPE_DISPATCHER,
			ComponentId:   zutil.IncSequence(),
			ListenAddr:    toAddr(ip, dispatcherPort+1),
			OwnerAddr:     center.ListenAddr,
			LogFile:       "dispatcher2.log",
			LogLevel:      "debug",
		},
	}
}
func setMongoAddr(config *FKVDBConfig, ip string) {
	config.Url = fmt.Sprintf("mongodb://%s:27017", ip)
}
func setServiceAddr(service *FServiceConfig, ip string, port int, ownerAddr string) {

	service.ListenAddr = fmt.Sprintf("%s:%d", ip, port)
	service.OwnerAddr = ownerAddr
}
func setServiceListAddr(serviceList []FServiceConfig, ip string, port int, ownerAddr string) {
	for k := range serviceList {
		serviceList[k].ListenAddr = fmt.Sprintf("%s:%d", ip, port+k)
		serviceList[k].OwnerAddr = ownerAddr
	}
}
func init() {
	ip := "127.0.0.1"
	ipServer := "124.221.147.27"
	ipClient := ip
	centerI := 0

	mongodb := initDB()
	version := initVersion(ip)
	loginList := initLoginList(ip, version)
	centerList := initCenterList(ip, version)
	center := &(centerList)[centerI]
	gateList := initGateList(ip, center)
	dispatcherList := initDispatcherList(ip, center)

	{
		serviceConfigFile := &FServiceConfigFile{
			Version:        *version,
			CenterList:     centerList,
			GateList:       gateList,
			LoginList:      loginList,
			DispatcherList: dispatcherList,
			MongoDB:        *mongodb,
		}
		bytes, _ := json.MarshalIndent(serviceConfigFile, "", "\t")
		file := filepath.Join(base.AppPath, "config/service.json")
		os.WriteFile(file, bytes, 0666)
	}

	{
		serviceConfigFile := &FServiceConfigFile{
			Version:   *version,
			LoginList: loginList,
			MongoDB:   *mongodb,
		}
		bytes, _ := json.MarshalIndent(serviceConfigFile, "", "\t")
		file := filepath.Join(base.AppPath, "config/service_version.json")
		os.WriteFile(file, bytes, 0666)
	}
	{
		ip = ipClient
		setServiceAddr(version, ipServer, versionPort, "")
		setServiceListAddr(centerList, ip, centerPort, version.ListenAddr)
		setServiceListAddr(gateList, ip, gatePort, center.ListenAddr)
		setServiceListAddr(dispatcherList, ip, dispatcherPort, center.ListenAddr)
		serviceConfigFile := &FServiceConfigFile{
			CenterList:     centerList,
			GateList:       gateList,
			DispatcherList: dispatcherList,
			MongoDB:        *mongodb,
		}
		bytes, _ := json.MarshalIndent(serviceConfigFile, "", "\t")
		file := filepath.Join(base.AppPath, "config/service_center.json")
		os.WriteFile(file, bytes, 0666)
	}
}

func main() {

}
