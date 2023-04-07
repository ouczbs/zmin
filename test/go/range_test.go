package _go

import (
	"github.com/ouczbs/zmin/engine/core/zlog"
	"testing"
)

func TestRangeIntArray(t *testing.T) {
	zlog.Debug(&intArray[0])
	RangeIntArray(intArray)
	RangeIntArrayPtr(&intArray)
}
func TestRangeStringSlice(t *testing.T) {
	zlog.Debug(&intList[0])
	RangeIntSlice(intList)
}
func TestRangeStructSlice(t *testing.T) {
	zlog.Debug(&structList[0])
	RangeStructSlice(structList)
}

func TestRangeStructMap(t *testing.T) {
	zlog.Debug(structMap[1])
	RangeStructMap(structMap)
}
