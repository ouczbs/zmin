package base

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"runtime"
)

type FServiceConfigFile struct {
	Version        FServiceConfig
	CenterList     []FServiceConfig
	LoginList      []FServiceConfig
	GateList       []FServiceConfig
	DispatcherList []FServiceConfig

	MongoDB FKVDBConfig
}

var (
	ServiceConfigFile FServiceConfigFile
	AppPath           string
	ComponentId       int32
)

func LoadServiceFile(serviceFile string) {
	_, curPath, _, ok := runtime.Caller(0)
	if ok {
		AppPath = filepath.Join(curPath, "../../")
	}
	appConfigPath := filepath.Join(AppPath, "config", serviceFile)
	bytes, _ := os.ReadFile(appConfigPath)
	json.Unmarshal(bytes, &ServiceConfigFile)
}
func init() {
	serviceFile := flag.String("config", "service.json", "config")
	cid := flag.Int("ComponentId", 0, "component id")
	flag.Parse()
	ComponentId = int32(*cid)
	LoadServiceFile(*serviceFile)
}
func getServiceConfig(ComponentId int32) *FServiceConfig {
	if ServiceConfigFile.Version.ComponentId == ComponentId {
		return &ServiceConfigFile.Version
	}
	for _, service := range ServiceConfigFile.CenterList {
		if service.ComponentId == ComponentId {
			return &service
		}
	}

	for _, service := range ServiceConfigFile.LoginList {
		if service.ComponentId == ComponentId {
			return &service
		}
	}

	for _, service := range ServiceConfigFile.GateList {
		if service.ComponentId == ComponentId {
			return &service
		}
	}

	for _, service := range ServiceConfigFile.DispatcherList {
		if service.ComponentId == ComponentId {
			return &service
		}
	}
	return nil
}
