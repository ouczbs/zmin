package zconf

type IModel interface {
	Table()string
	M()map[string]interface{}
}