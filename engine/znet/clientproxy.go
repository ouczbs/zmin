package znet

import (
	"Zmin/engine/zclass"
	"Zmin/engine/zlog"
	"net"
)


type UClientProxy struct {
	*UPacketConnection
	*UProperty
	owner    IService

	ReqHandleMaps map[TCallId]FRequestHandle

}

func NewClientProxy(owner IService, conn net.Conn) *UClientProxy {
	packetConn := NewPacketConnection(conn)
	proxy := &UClientProxy{
		UPacketConnection:packetConn,
		owner : owner,
		UProperty:zclass.NewProperty(),
		ReqHandleMaps:make(map[TCallId]FRequestHandle),
	}
	return proxy
}
func (proxy * UClientProxy) ClientDisconnect(){
	if err := recover(); err != nil {
		zlog.Errorf("Client %s paniced with error: %v", proxy, err)
	}
	proxy.Close()
	proxy.owner.ClientDisconnect(proxy)
}
func (proxy *UClientProxy) Serve() {
	// Serve the dispatcher client from server / gate
	for {
		packet, err := proxy.RecvPacket()
		if err != nil {
			packet.Release()
			proxy.ClientDisconnect()
			zlog.Debugf("Serve RecvPacket error ",err  )
			return
		}
		messageType := packet.ReadMessageType()
		message := NewMessage(proxy, messageType , packet )
		proxy.owner.RecvMessage(message)
	}
}
func (proxy *UClientProxy) GetRequestHandles(id TCallId , cmd TCmd)(FRequestHandle , FRequestHandle){
	handle := proxy.ReqHandleMaps[id]
	if handle != nil {
		delete(proxy.ReqHandleMaps , id)
	}
	globalHandle := proxy.owner.GetRequestHandle(cmd)
	return handle,globalHandle
}
func (proxy *UClientProxy) Then(handle FRequestHandle , request * URequest) *UClientProxy{
	if handle == nil {
		return proxy
	}
	if request.Next {
		request.Next = false
		handle(proxy, request)
	}
	return proxy
}
