package entities

type person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) person {
	var person = person{name, age}

	return person
}
