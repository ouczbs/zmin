package string

import (
	"fmt"
	"reflect"
	"unsafe"
)

func String2Bytes(s *string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) *string {
	return (*string)(unsafe.Pointer(&b))
}
func String2Bytes2(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String2(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func StringAddr(s string) string {
	fmt.Printf("recv string %p \n", &s)
	return s
}
