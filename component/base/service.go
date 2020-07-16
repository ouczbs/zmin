package base

import (
	"Zmin/engine/zattr"
	"Zmin/engine/zclass"
	"Zmin/engine/zconf"
	"Zmin/engine/zlog"
	"Zmin/engine/znet"
	"Zmin/engine/zproto/pb"
	"net"
	"os"
)

type UService struct {
	*UProperty

	ReqHandleMaps 		map[TCmd]FRequestHandle
	MessageQueue  		chan * UMessage
}
func NewService(reqHandleMaps map[TCmd]FRequestHandle)*UService{
	return &UService{
		zclass.NewProperty(),
		reqHandleMaps,
		make(chan *UMessage, zconf.CQueueMessageSize),
	}
}
func (service * UService) ClientDisconnect(proxy *UClientProxy) {
	zlog.Debugf("ClientDisconnect %s" ,proxy)
	t ,ok := proxy.GetProperty(zattr.Uint32ComponentType).(uint32)
	if ok &&  t  == uint32(pb.COMPONENT_TYPE_CENTER){
		zlog.Infof(" Center client notify exit process !!!")
		service.Close()
	}
}
func (service * UService)Close() {
	os.Exit(1)
}
func (service * UService)NewTcpConnection(conn net.Conn) {
	proxy := znet.NewClientProxy(service , conn)
	proxy.Serve()
}
func (service * UService) RecvMessage(message *UMessage) {
	service.MessageQueue <- message
}
func (service * UService) GetRequestHandle(cmd TCmd) FRequestHandle {
	return service.ReqHandleMaps[cmd]
}
func (service *UService) SyncProxyProperty(proxy *UClientProxy, request * URequest){
	message , ok := request.ProtoMessage.(* pb.SYNC_PROXY_PROPERTY)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : " , proxy , request)
		return
	}
	for _,property := range message.PropertyMapList{
		proxy.SetProperty(property.Key ,zattr.ReadProperty(property.Value))
		zlog.Debug("SyncProxyProperty:attr k:" , property.Key  )
	}
}
func (service *UService) InitDownHandles(){
	service.SetProperty(zattr.StringCenterAddr , zconf.GetCenterConfig().ListenAddr)
	service.ReqHandleMaps[TCmd(pb.CommandList_MT_SYNC_PROXY_PROPERTY)] = service.SyncProxyProperty
}
func (service *UService)MakeClientProxy(addr string)*znet.UClientProxy{
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		zlog.Infof(" MakeClientProxy error , addr %s , err %s" ,addr , err)
		return nil
	}
	proxy := znet.NewClientProxy(service , conn)
	go proxy.Serve()
	return proxy
}
func (service *UService)MakeCenterProxy()*znet.UClientProxy{
	addr ,ok := service.GetProperty(zattr.StringCenterAddr).(string)
	if !ok {
		zlog.Error("ConnectToCenter :attr k:" , zattr.StringListenAddr)
		return nil
	}
	centerProxy := service.MakeClientProxy(addr)
	return centerProxy
}