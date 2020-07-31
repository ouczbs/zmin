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
	MT_SYNC_PROPERTY
	MT_BROADCAST
	MT_TO_CLIENT
	MT_TO_GATE
	MT_TO_CENTER
	MT_FROM_CENTER
	MT_TO_LOGIN
	MT_TO_GAME_START = 100 + iota

	MT_TO_GAME_END = 1000 + iota

	CodeOk = iota
	CodeError
)
