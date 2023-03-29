package zmodel

import (
	"github.com/ouczbs/zmin/engine/data/zconf"
)

type (
	IModel = zconf.IModel

	TComponentType = zconf.TComponentType
	TComponentId   = zconf.TComponentId
)

type FEncoder func(field uintptr) interface{}
type Field struct {
	name      string
	offset    uintptr
	sliceSize int
	encoder   FEncoder
}
