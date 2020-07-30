package zclass

import (
	"github.com/ouczbs/Zmin/engine/zconf"
	"sync"
)

type TEnum = zconf.TEnum

type UProperty struct {
	propertyLock sync.RWMutex
	Property map[TEnum]interface{}
}

func NewProperty()*UProperty{
	return &UProperty{
		Property:make(map[TEnum]interface{}),
	}
}
//设置链接属性
func (pro *UProperty) SetProperty(key TEnum, value interface{}) {
	pro.propertyLock.Lock()
	pro.Property[key] = value
	pro.propertyLock.Unlock()
}
//获取链接属性
func (pro *UProperty) GetProperty(key TEnum) interface{}{
	pro.propertyLock.RLock()
	value, _ := pro.Property[key]
	pro.propertyLock.RUnlock()
	return value
}
//移除链接属性
func (pro *UProperty) RemoveProperty(key TEnum) {
	pro.propertyLock.Lock()
	delete(pro.Property, key)
	pro.propertyLock.Unlock()
}
