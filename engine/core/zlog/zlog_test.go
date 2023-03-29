package zlog

import (
	"go.uber.org/zap"
	"testing"
)

func TestDebugf(t *testing.T) {
	//SetSource("zlog_test")
	SetOutput([]string{"stderr", "zlog_test.log"})
	Debugf("test sugar.Debugf %d", 1)
	Error("test sugar.Error ", 2)
	sugar.Errorf("test sugar.Error %d", 2)
	logger.Debug("test logger.Debug" , zap.Int("logger",3))
}

