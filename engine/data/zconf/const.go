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
	COMPONENT_TYPE_CENTER     = iota
	COMPONENT_TYPE_DISPATCHER = 1
	COMPONENT_TYPE_LOGIN      = 2
	COMPONENT_TYPE_GATE       = 3
	COMPONENT_TYPE_GAME       = 4
)
