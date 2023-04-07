package _go

import (
	"github.com/ouczbs/zmin/engine/core/zlog"
)

func RangeIntArray(array [10]int) {
	for k, v := range array {
		zlog.Debug(&k, &v, &array[k])
	}
}
func RangeIntArrayPtr(array *[10]int) {
	for k, v := range array {
		zlog.Debug(&k, &v, &array[k])
	}
}

func RangeIntSlice(array []int) {
	for k, v := range array {
		zlog.Debug(&k, &v, &array[k])
	}
}
func RangeStructSlice(array []arrayStruct) {
	for k, v := range array {
		zlog.Debug("++++++++++++++++++")
		zlog.Debug(&v.id, &array[k].id)
		zlog.Debug(&v.name, &array[k].name)
		zlog.Debug(&k, &v, &array[k])
	}
}

func RangeStructMap(array map[int]arrayStruct) {
	for k, v := range array {
		zlog.Debug("++++++++++++++++++")
		d := v.data
		d.id = k * 10
		vp := v.parent
		vp.id = k * 10
		vp.name = vp.name + "xxx"
		zlog.Debug(vp.id, array[k].parent.id)
		zlog.Debug(d.id, v.data.id, array[k].data.id)
		zlog.Debug(k, &d.id, &vp.data.id, array[k].data)
	}
}
