package zconf

const (
	CQueuePacketSize  = 1024
	CQueueRequestSize = 10000
	CQueueProxySize   = 10000
	CQueueMessageSize = 10000

	CPoolPacketSize        = CQueuePacketSize
	CPoolMessageSize       = CQueuePacketSize
	CPoolRequestSize       = CQueuePacketSize
	CMinPacketBuffer       = 128
	CMaxPacketBuffer       = 1024 * 1024 * 16
	CPacketHeadSize        = 2
	CPacketMessageTypeSize = 2
	CPacketRequestTypeSize = 2
)
const (
	COMPONENT_TYPE_VERSION = iota
	COMPONENT_TYPE_CENTER
	COMPONENT_TYPE_DISPATCHER
	COMPONENT_TYPE_LOGIN
	COMPONENT_TYPE_GATE
	COMPONENT_TYPE_GAME
)
