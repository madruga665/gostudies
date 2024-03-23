package entities

import "testing"

func TestNewPerson(t *testing.T) {
	nameExpectedResult := "Luciano"
	ageExpectedResult := 37
	person := NewPerson("Luciano", 37)

	if person.Name != nameExpectedResult {
		t.Errorf("O nome era pra ser %s, mas saiu %s", nameExpectedResult, person.Name)
	}

	if person.Age != ageExpectedResult {
		t.Errorf("O nome era pra ser %d, mas saiu %d", ageExpectedResult, person.Age)
	}
}
