package gate

import (
	"zmin/engine/zutil"
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/zattr"
	"github.com/ouczbs/zmin/engine/zconf"
	"github.com/ouczbs/zmin/engine/zlog"
	"github.com/ouczbs/zmin/engine/znet"
	"net"
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
func (service *UGateService) NewTcpConnection(conn net.Conn) {
	proxy := znet.NewClientProxy(service, conn)
	id := zutil.IncSequence()
	clientProxyMaps[id] = proxy
	proxy.SetProperty(zattr.Int32ComponentId, int32(id))
	proxy.Serve()
}