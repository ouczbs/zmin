package zconf

const (
	CMD_PROXY_CMD = iota * 256 
	CMD_LOGIN_CMD = 2 * 256
)

const (
	CMD_LoginAccount = CMD_LOGIN_CMD + iota
	CMD_LoginAccountAck 
	CMD_RegisterAccount 
)
const (
	CMD_ADD_ENGINE_COMPONENT = CMD_PROXY_CMD + iota
	CMD_ADD_ENGINE_COMPONENT_ACK 
	CMD_SYNC_PROXY_PROPERTY 
	CMD_SET_REMOTE_PROPERTY 
	CMD_ADD_CLIENT_ACK 
)

var (
	CommandList_name = map[int32]string{
		CMD_LoginAccount : "zpb.LoginAccount",
		CMD_LoginAccountAck : "zpb.LoginAccountAck",
		CMD_RegisterAccount : "zpb.RegisterAccount",
		CMD_ADD_ENGINE_COMPONENT : "zpb.ADD_ENGINE_COMPONENT",
		CMD_ADD_ENGINE_COMPONENT_ACK : "zpb.ADD_ENGINE_COMPONENT_ACK",
		CMD_SYNC_PROXY_PROPERTY : "zpb.SYNC_PROXY_PROPERTY",
		CMD_SET_REMOTE_PROPERTY : "zpb.SET_REMOTE_PROPERTY",
		CMD_ADD_CLIENT_ACK : "zpb.ADD_CLIENT_ACK",
	}
)
