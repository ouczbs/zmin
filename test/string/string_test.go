package string

import (
	"fmt"
	"testing"
)

//	type stringStruct struct {
//		str unsafe.Pointer
//		len int
//	}
var (
	testStr = []byte("Hello Gopher! Hello Gopher! Hello Gopher!")
)

func BenchmarkBytes2String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := Bytes2String(testStr)
		String2Bytes(str)
	}
}
func BenchmarkBytes2String2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := Bytes2String2(testStr)
		String2Bytes2(str)
	}
}
func TestStringAddr(t *testing.T) {
	str := Bytes2String2(testStr)
	res := StringAddr(str)
	fmt.Printf("sendstr: %p returnstr: %p \n", &str, &res)
}
