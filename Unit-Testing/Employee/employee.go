package employee

type emp struct {
	name string
	age  int
}

func checkAge(e emp) (bool, emp) {
	if e.age < 22 {
		return false, emp{}
	}

	return true, e
}
