package entities

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	person := Person{name, age}

	return person
}
