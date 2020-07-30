package dispatcher

import (
	"github.com/ouczbs/Zmin/component/base"
	"github.com/ouczbs/Zmin/engine/zattr"
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zlog"
	"github.com/ouczbs/Zmin/engine/znet"
)

type UDispatcherService struct {
	*UService
}
func NewDispatcherService() *UDispatcherService {
	return &UDispatcherService{
		UService:base.NewService(reqHandleMaps),
	}
}
func (service * UDispatcherService)Run() {
	service.UService.Run()
	service.initService()
	go service.MessageLoop()
	service.ConnectToCenter()
	znet.ServeTCPForever(service.Config.ListenAddr, service)
}
func (service *UDispatcherService) initService(){
	logFile,ok := service.GetProperty(zattr.StringLogFile).(string)
	if !ok{
		logFile = zconf.DispatcherConfig.LogFile
	}
	zlog.SetOutput([]string{ "stderr", logFile })
	service.InitDownHandles()
}