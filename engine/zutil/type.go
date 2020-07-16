package zutil

import "Zmin/engine/zconf"

type (
	TSequence = zconf.TSequence
)

var (
	sequence TSequence = 0
)

type timeoutError interface {
	Timeout() bool // Is it a timeout error
}