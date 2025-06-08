package main

import (
	"testing"
)

func TestCalculoMedia_ValoresValidos(t *testing.T) {

	resultado := calculoMedia(5, 10, 7, 8, 9, 6, 4)
	esperado := 7
	if resultado != esperado {
		t.Errorf("Esperado %d, mas obteve %d", esperado, resultado)
	}
}

func TestCalculoMedia_Negativo(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Era esperado panic ao receber número negativo")
		}
	}()
	_ = calculoMedia(5, 4, -1, 8)
}

func TestCalculoMedia_SoUmValor(t *testing.T) {
	resultado := calculoMedia(10)
	esperado := 10
	if resultado != esperado {
		t.Errorf("Esperado %d, mas obteve %d", esperado, resultado)
	}
}
