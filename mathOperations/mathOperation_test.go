package mathOperations

import "testing"

func TestSum(t *testing.T) {
	want := 6
	sum := Sum(3, 3)
	if want != sum {
		t.Fatalf("O resultado era pra ser %d mas foi %d", want, sum)
	}

}

func TestSubtract(t *testing.T) {
	want := 2
	subtract := Subtract(5, 3)
	if want != subtract {
		t.Errorf("O resultado era pra ser %d mas foi %d", want, subtract)
	}

}

func TestMultiply(t *testing.T) {
	want := 15
	multiply := Multiply(5, 3)
	if want != multiply {
		t.Errorf("O resultado era pra ser %d mas foi %d", want, multiply)
	}

}

func TestDivision(t *testing.T) {
	want := 10
	division := Division(20, 2)

	if want != division {
		t.Errorf("O resultado era pra ser %d mas foi %d", want, division)
	}
}
