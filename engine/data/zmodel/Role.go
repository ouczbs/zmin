package zmodel

type FRole struct {
	Id       int32
	OwnerID  int32
	ServerID int32
	SceneID  int32
	GateID   int32
	Player   string
}

func (role *FRole) Table() string {
	return "role"
}

func (role *FRole) M() map[string]interface{} {
	return M(role)
}

var Role = &FRole{}

func init() {
	Schema(Role)
}
