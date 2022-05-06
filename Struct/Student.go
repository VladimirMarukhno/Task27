package Struct

type Student struct {
	Name  string
	Age   int
	Grade int
}

func (s Student) Put() string {
	return s.Name
}
