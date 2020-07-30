package zmodel

import (
	"github.com/ouczbs/Zmin/engine/zconf"
	"sync"
	"unsafe"
)
type IModel = zconf.IModel
var dbCache sync.Map
type FEncoder func(field uintptr)interface{}
type Field struct {
	name string
	offset uintptr
	sliceSize int
	encoder FEncoder
}
type emptyInterface struct {
	typ  *struct{}
	word unsafe.Pointer
}
