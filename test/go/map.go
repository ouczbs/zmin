package _go

var (
	structMap map[int]arrayStruct
)

func init() {
	structMap = make(map[int]arrayStruct)
	p1 := &arrayStruct{
		id:   1,
		name: "p1",
	}
	p2 := &arrayStruct{
		id:   2,
		name: "p1",
	}
	p3 := &arrayStruct{
		id:   3,
		name: "p1",
	}
	d1 := arrayData{
		id: 1,
	}
	d2 := arrayData{
		id: 2,
	}
	d3 := arrayData{
		id: 3,
	}
	structMap[1] = arrayStruct{
		id:     1,
		name:   "name1",
		parent: p1,
		data:   d1,
	}
	structMap[2] = arrayStruct{
		id:     2,
		name:   "name2",
		parent: p2,
		data:   d2,
	}
	structMap[3] = arrayStruct{
		id:     3,
		name:   "name3",
		parent: p3,
		data:   d3,
	}
}
