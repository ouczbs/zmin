package znet

import (
	"github.com/ouczbs/zmin/engine/zlog"
	"net"
)

type UPacketConnection struct {
	Conn net.Conn
	isClosed bool
}

//创建连接的方法
func NewPacketConnection(conn net.Conn) *UPacketConnection {
	//初始化Conn属性
	c := &UPacketConnection{
		Conn:         conn,
		isClosed:     false,
	}
	return c
}

func (pc *UPacketConnection) read(bytes []byte , s TSize, e TSize) error{
	n := int(e - s)
	rn := 0
	for rn < n {
		cn, err := pc.Conn.Read(bytes[s:e])
		if err != nil {
			zlog.Debug("Conn read error: " , err)
			return err
		}
		rn += cn
	}
	return nil
}
/*
	读消息Goroutine，用于从客户端中读取数据
*/
func (pc *UPacketConnection) RecvPacket()(* UPacket , error){
	packet := NewPacket()
	err := pc.read(packet.bytes , 0 , _CPacketHeadSize)
	if err != nil {
		return packet, err
	}
	size := packet.ReadSize()
	packet.SetSize(size)
	err = pc.read(packet.bytes , _CPacketHeadSize , size)
	return packet , err
}
func (pc *UPacketConnection) SendPacket(packet * UPacket)error{
	size := int(packet.Size)
	bytes := packet.bytes
	packet.WriteSize(packet.Size)
	for size > 0 {
		n, err := pc.Conn.Write(bytes[:size])
		if n == size && err == nil { // handle most common case first
			return nil
		}
		if n > 0 {
			bytes = bytes[n:]
			size -= n
		}
		if err != nil  {
			zlog.Debug("Conn write error: " , err)
			return err
		}
	}
	return nil
}
func (pc *UPacketConnection) ForwardPacket(packet * UPacket)error{
	err := pc.SendPacket(packet)
	packet.Release()
	return err
}
func (pc *UPacketConnection) Close() {
	pc.isClosed = true
}