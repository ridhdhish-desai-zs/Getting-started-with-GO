package employee

type emp struct {
	id     int
	name   string
	hasPan bool
	age    int
}

var employees = []emp{
	emp{id: 1, name: "Naruto", hasPan: false, age: 17},
	emp{id: 2, name: "Pain", hasPan: true, age: 24},
}

func (e *emp) setAge(value int) {
	e.age = value
}

func (e *emp) setName(value string) {
	e.name = value
}

func (e *emp) setHasPan(value bool) {
	e.hasPan = value
}

func (e *emp) setId(value int) {
	e.id = value
}

func checkAge(e emp) (bool, emp) {
	if e.age < 22 {
		return false, emp{}
	}

	return true, e
}

func getDetails(value int) emp {
	for _, v := range employees {
		if v.id == value {
			return v
		}
	}
	return emp{}
}
