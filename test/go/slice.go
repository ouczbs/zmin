package _go

var (
	intList    []int
	stringList []string
	structList []arrayStruct
	dataList   []arrayData
)

func init() {
	intList = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	structList = []arrayStruct{
		{
			id:   1,
			name: "name1",
		},
		{
			id:   2,
			name: "name2",
		},
		{
			id:   3,
			name: "name3",
		},
	}
}
