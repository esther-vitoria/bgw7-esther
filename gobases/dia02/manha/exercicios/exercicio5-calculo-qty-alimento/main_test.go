package main

import (
	"testing"
)

func floatEquals(a, b float64) bool {
	const eps = 0.0001
	if a > b {
		return a-b < eps
	}
	return b-a < eps
}

func TestQtyDog(t *testing.T) {
	result := qtyDog(3)
	expected := 30.0
	if !floatEquals(result, expected) {
		t.Errorf("qtyDog(3) = %.2f, esperado %.2f", result, expected)
	}
}

func TestQtyCat(t *testing.T) {
	result := qtyCat(2)
	expected := 10.0
	if !floatEquals(result, expected) {
		t.Errorf("qtyCat(2) = %.2f, esperado %.2f", result, expected)
	}
}

func TestQtyHamster(t *testing.T) {
	result := qtyHamster(8)
	expected := 2.0 // 8 * 0.25
	if !floatEquals(result, expected) {
		t.Errorf("qtyHamster(8) = %.2f, esperado %.2f", result, expected)
	}
}

func TestQtyTarantula(t *testing.T) {
	result := qtyTarantula(4)
	expected := 0.6 // 4 * 0.15
	if !floatEquals(result, expected) {
		t.Errorf("qtyTarantula(4) = %.2f, esperado %.2f", result, expected)
	}
}

func TestAnimalFunc(t *testing.T) {
	tests := []struct {
		animal    string
		qty, want float64
	}{
		{dog, 7, 70.0},
		{cat, 1, 5.0},
		{hamster, 10, 2.5},
		{tarantula, 2, 0.3},
	}
	for _, test := range tests {
		fn, err := Animal(test.animal)
		if err != nil {
			t.Errorf("Animal(%s): erro inesperado: %v", test.animal, err)
			continue
		}
		result := fn(test.qty)
		if !floatEquals(result, test.want) {
			t.Errorf("Animal(%s) com qty %.2f = %.2f, esperado %.2f", test.animal, test.qty, result, test.want)
		}
	}
}

func TestAnimalInvalid(t *testing.T) {
	_, err := Animal("elephant")
	if err == nil {
		t.Errorf("Esperado erro para animal inválido, obteve nil")
	}
}
