package zmodel

type FAccount struct {
	ID       int32
	UserName string
	Password string
	Phone    string
	Email    string
}

func (service *FAccount) Table() string {
	return "account"
}

func (service *FAccount) M() map[string]interface{} {
	return M(service)
}

var Account = &FAccount{}

func init() {
	Schema(Account)
}
