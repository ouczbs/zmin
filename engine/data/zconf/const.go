package zconf

const (

	CQueuePacketSize  = 1024
	CQueueRequestSize = 10000
	CQueueProxySize = 10000
	CQueueMessageSize = 10000

	CPoolPacketSize = CQueuePacketSize
	CPoolMessageSize = CQueuePacketSize
	CPoolRequestSize = CQueuePacketSize
	CMinPacketBuffer = 128
	CMaxPacketBuffer = 1024 * 1024 * 16
	CPacketHeadSize = 2
	CPacketMessageTypeSize = 2
	CPacketRequestTypeSize = 2
)
const (
	MT_INVALID TMessageType = iota

	MT_TO_ALL
	MT_TO_SERVER
	MT_TO_CLIENT
	MT_BROADCAST


	MT_TO_GAME_START = 1000

	MT_TO_GAME_END = 10000

	CodeOk = 0
	CodeError
)
