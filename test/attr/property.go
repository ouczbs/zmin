package attr

import (
	"Zmin/engine/zattr"
	"Zmin/engine/zlog"
	"encoding/json"
	"sync"
)

type UProperty struct {
	propertyLock sync.RWMutex
	property map[TEnum]interface{}
}
func (p *UProperty) SetProperty(key TEnum, value interface{}) {
	p.propertyLock.Lock()
	p.property[key] = value
	p.propertyLock.Unlock()
	zlog.Debug(&value)
}
//获取链接属性
func (p *UProperty) GetProperty(key TEnum) interface{}{
	p.propertyLock.RLock()
	value, ok := p.property[key]
	p.propertyLock.RUnlock()
	if !ok {
		zlog.Infof("GetProperty error , key = %s " , key)
	}
	return value
}

func Marsh(m map[TEnum]interface{}){
	c, _ := json.Marshal(m)
	zlog.Infof( string(c))
	UnMarsh(c)
}
func UnMarsh(bytes []byte){
	p := &UProperty{}
	p.property = make(map[TEnum]interface{})
	err := json.Unmarshal(bytes ,&p.property  )
	if err != nil {
		zlog.Error(err)
	}
	zlog.Debug(p.property )
	zlog.Debug(p.GetProperty(zattr.StringListenAddr).(string))
}