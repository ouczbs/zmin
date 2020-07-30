package attr

import (
	"Zmin/engine/zattr"
	"Zmin/engine/zlog"
	"testing"
)

func TestUProperty_SetProperty(t *testing.T) {
	p := &UProperty{}
	p.property = make(map[TEnum]interface{})
	v := uint32(2)
	p.SetProperty(zattr.Int32ComponentId, v)
	zlog.Debug(&v)
	id, ok := p.GetProperty(zattr.Int32ComponentType).(int32)
	zlog.Debug(id, ok)
	s := "dfsdgffg"
	b := []byte(s)
	b[0] = 1
	convert(s)
}
func convert(v interface{}) {
	zlog.Debug(v.([]byte))
}

func TestMarsh(t *testing.T) {
	p := &UProperty{}
	p.property = make(map[TEnum]interface{})
	p.SetProperty(zattr.StringListenAddr, string("127.0.0.1:11111"))
	p.SetProperty(zattr.BoolIsLoadedService, false)
	p.SetProperty(zattr.Int32ComponentId, int32(2))

	b1 := "127.0.0.1:"
	b2 := ":11111"
	b3 := b1 + b2
	zlog.Debug(b1, b2, b3)
	//Marsh(p.property)

}
