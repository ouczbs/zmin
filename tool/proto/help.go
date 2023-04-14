package main

import (
	"flag"
	"github.com/ouczbs/zmin/engine/core/zlog"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"unsafe"
)

func MakeAppPath() string {
	var AppPath string
	_, curPath, _, ok := runtime.Caller(0)
	if ok {
		AppPath = filepath.Join(curPath, "../../../")
	}
	return AppPath
}
func GetProtoFiles(AppPath string) []os.DirEntry {
	in := flag.String("proto", INDIR, "proto")
	out := flag.String("out", OUTFILE, "out")
	flag.Parse()
	outFile = filepath.Join(AppPath, *out)
	inDir = filepath.Join(AppPath, *in)
	files, err := os.ReadDir(inDir)
	if err != nil {
		zlog.Debugf("error %s \n", err)
		return nil
	}
	return files
}
func ReadFile(file string) string {
	_bytes, err := os.ReadFile(file)
	if err != nil {
		zlog.Debugf("error %s \n", err)
	}
	return Bytes2String(_bytes)
}
func FlushFile() {
	os.WriteFile(outFile, String2Bytes(bytes.String()), 0666)
	bytes.Reset()
}
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
