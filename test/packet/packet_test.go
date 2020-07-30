package packet

import (
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zlog"
	"testing"
	"time"
)

var packetList = make(chan *UPacket , zconf.CPoolPacketSize * 2)
func runNewPacket() {
	for true {
		packet := NewPacket()
		packetList <- packet
	}
}
func runRelease(){
	for true {
		packet := <- packetList
		packet.Release()
	}
}
func runPacketPool()  {
	for j := 0; j < 300 ; j++ {
		go runNewPacket()
		go runNewPacket()
		go runRelease()
	}
	time.Sleep(time.Duration(1) * time.Second / 3)
	zlog.Debugf("push count  = %d , pop count = %d" , pushCount , popCount)
}
func TestNewPacketPool(t *testing.T) {
	stackPoolType = 1
	runPacketPool()
}
func TestNewPacketStackPoolPool(t *testing.T) {
	stackPoolType = 2
	runPacketPool()
}
func getBytes()[]byte{
	bytes := make([]byte , 128)
	zlog.Debug(bytes, &bytes[0])
	return bytes
}
func TestUPacket_Init(t *testing.T) {
	bytes := getBytes()
	packet := NewPacket()
	packet.bytes = bytes
	zlog.Debug(packet.bytes, &packet.bytes[0])
}