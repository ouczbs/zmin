package zmodel

import (
	"Zmin/engine/zlog"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"strings"
	"unsafe"
)


func Schema(v IModel){
	key := v.Table()
	if _,ok := dbCache.Load(key); ok{
		return
	}
	e := reflect.TypeOf(v).Elem()
	value := make([]*Field , e.NumField())
	for i:=0 ;i<e.NumField();i++{
		f := e.Field(i)
		value[i] = &Field{
			name: strings.ToLower(f.Name),
			offset: f.Offset,
			encoder: NewTypeEncoder(f.Type),
		}
	}
	dbCache.LoadOrStore(key , value)
}

func M(v IModel)bson.M{
	value ,ok := dbCache.Load(v.Table())
	if !ok {
		return nil
	}
	structPtr := (*emptyInterface)(unsafe.Pointer(&v)).word
	fieldList := value.([]*Field)
	m := bson.M{}
	for _,field := range fieldList{
		ptr := uintptr(structPtr) + field.offset
		m[field.name] = field.encoder(ptr)
	}
	return m
}
func NewTypeEncoder(t reflect.Type)FEncoder{
	switch t.Kind() {
	case reflect.String:
		return stringEncoder
	case reflect.Int,reflect.Int32:
		return intEncoder
	case reflect.Uint,reflect.Uint32:
		return uintEncoder
	case reflect.Bool:
		return boolEncoder
	case reflect.Float32:
		return float32Encoder
	case reflect.Float64:
		return float64Encoder
	default:
		zlog.Error("Mongo Schema can‘t support this type : " , t.Kind())
		return defaultEncoder
	}
}
func defaultEncoder(_ uintptr)interface{}{
	return nil
}
func stringEncoder(field uintptr)interface{}{
	return *(*string)(unsafe.Pointer(field))
}
func boolEncoder(field uintptr)interface{}{
	return *(*bool)(unsafe.Pointer(field))
}
func intEncoder(field uintptr)interface{}{
	return *(*int32)(unsafe.Pointer(field))
}
func uintEncoder(field uintptr)interface{}{
	return *(*uint32)(unsafe.Pointer(field))
}
func float32Encoder(field uintptr)interface{}{
	return *(*float32)(unsafe.Pointer(field))
}
func float64Encoder(field uintptr)interface{}{
	return *(*float64)(unsafe.Pointer(field))
}