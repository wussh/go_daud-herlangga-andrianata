package calculation

import "testing"

func TestAdditionPositip(t *testing.T) {
	result := Additon(1, 2)
	if result != 3 {
		t.Errorf("Expect 3, get %d", result)
	}
}

func TestAdditionNegatip(t *testing.T) {
	result := Additon(-1, -2)
	if result != -3 {
		t.Errorf("Expect -3, get %d", result)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(1, 2)
	if result != -1 {
		t.Errorf("Expect -1, get %d", result)
	}
}
func TestDivision(t *testing.T) {
	result := Division(4, 2)
	if result != 2 {
		t.Errorf("Expect 2, get %d", result)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(4, 2)
	if result != 8 {
		t.Errorf("Expect 8, get %d", result)
	}
}
