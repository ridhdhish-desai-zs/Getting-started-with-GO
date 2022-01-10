package employee

type emp struct {
	id     int
	name   string
	hasPan bool
	age    int
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

func checkAge(e emp) (bool, emp) {
	if e.age < 22 {
		return false, emp{}
	}

	return true, e
}
