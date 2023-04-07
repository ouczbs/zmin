package _go

type arrayData struct {
	id int
}
type arrayStruct struct {
	id     int
	name   string
	parent *arrayStruct
	data   arrayData
}

var (
	intArray    [10]int
	stringArray [10]string
	structArray [10]arrayStruct
	dataArray   [10]arrayData
)
