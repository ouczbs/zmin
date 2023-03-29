package zutil

import (
	"github.com/ouczbs/zmin/engine/core/zbase"
)

type (
	TSequence = zbase.TSequence
)

type timeoutError interface {
	Timeout() bool // Is it a timeout error
}
