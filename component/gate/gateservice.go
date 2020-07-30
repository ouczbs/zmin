package gate

import (
	"github.com/ouczbs/Zmin/component/base"
	"github.com/ouczbs/Zmin/engine/zattr"
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zlog"
	"github.com/ouczbs/Zmin/engine/znet"
)

type UGateService struct {
	*UService
}

func NewGateService() *UGateService {
	return &UGateService{
		UService:base.NewService(reqHandleMaps),
	}
}
func (service * UGateService)Run() {
	service.UService.Run()
	service.initService()
	go service.MessageLoop()
	service.ConnectToCenter()
	znet.ServeTCPForever(service.Config.ListenAddr, service)
}
func (service *UGateService) initService(){
	logFile,ok := service.GetProperty(zattr.StringLogFile).(string)
	if !ok{
		logFile = zconf.GateConfig.LogFile
	}
	zlog.SetOutput([]string{ "stderr", logFile })
	service.InitDownHandles()
}