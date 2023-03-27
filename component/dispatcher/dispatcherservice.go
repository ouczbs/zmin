package dispatcher

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/zattr"
	"github.com/ouczbs/zmin/engine/zconf"
	"github.com/ouczbs/zmin/engine/zlog"
	"github.com/ouczbs/zmin/engine/znet"
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