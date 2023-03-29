package zmodel

type FService struct {
	Id         TComponentId
	Type       TComponentType
	ListenAddr string
	Path       string
	Property   string
}

func (service *FService) Table() string {
	return "service"
}

func (service *FService) M() map[string]interface{} {
	return M(service)
}

var Service = &FService{}

func init() {
	Schema(Service)
}
