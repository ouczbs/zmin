package gate

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/core/zutil"
	"github.com/ouczbs/zmin/engine/net/znet"
	"github.com/ouczbs/zmin/engine/sync/zattr"
	"net"
)

type UGateService struct {
	*UService
}

func NewGateService() *UGateService {
	return &UGateService{
		UService: base.NewService(reqHandleMaps),
	}
}
func (service *UGateService) Run() {
	service.UService.Run()
	service.initService()
	go service.MessageLoop()
	service.ConnectToCenter()
	znet.ServeTCPForever(service.Config.ListenAddr, service)
}
func (service *UGateService) initService() {
	logFile, ok := service.GetProperty(zattr.StringLogFile).(string)
	if !ok {
		logFile = base.GateConfig.LogFile
	}
	zlog.SetOutput([]string{"stderr", logFile})
	service.InitDownHandles()
}
func (service *UGateService) NewTcpConnection(conn net.Conn) {
	proxy := znet.NewClientProxy(service, conn)
	id := zutil.IncSequence()
	clientProxyMaps[id] = proxy
	proxy.SetProperty(zattr.Int32ComponentId, int32(id))
	proxy.Serve()
}
