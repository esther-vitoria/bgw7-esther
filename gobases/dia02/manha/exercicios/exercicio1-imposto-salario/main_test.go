package main

import (
	"testing"
)

func TestCalcularImposto_AbaixoDe50000(t *testing.T) {
	salario := 45000.0
	esperado := salario * 0.10 // 10%
	resultado := impostoSalario(salario)
	if resultado != esperado {
		t.Errorf("Esperado %.2f, mas obteve %.2f", esperado, resultado)
	}
}

func TestCalcularImposto_AcimaDe50000(t *testing.T) {
	salario := 149999.3
	esperado := salario * 0.17 // 17%
	resultado := impostoSalario(salario)
	if resultado != esperado {
		t.Errorf("Esperado %.2f, mas obteve %.2f", esperado, resultado)
	}
}

func TestCalcularImposto_AcimaDe150000(t *testing.T) {
	salario := 200000.0
	esperado := salario * 0.27 // 27%
	resultado := impostoSalario(salario)
	if resultado != esperado {
		t.Errorf("Esperado %.2f, mas obteve %.2f", esperado, resultado)
	}
}
