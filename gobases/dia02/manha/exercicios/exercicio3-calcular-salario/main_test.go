package main

import (
	"math"
	"testing"
)

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 0.00001
}

func TestCalculoSalario_CategoriaC(t *testing.T) {
	minutos := 120.0 // 2 horas
	esperado := 2 * 1000.0
	resultado := calculoSalario(minutos, categoriaC)
	if !floatEquals(resultado, esperado) {
		t.Errorf("Esperado %.2f, obteve %.2f", esperado, resultado)
	}
}

func TestCalculoSalario_CategoriaB(t *testing.T) {
	minutos := 180.0                              // 3 horas
	salarioBase := 3 * 1500.0                     // 4500
	esperado := salarioBase + (salarioBase * 0.2) // 5400
	resultado := calculoSalario(minutos, categoriaB)
	if !floatEquals(resultado, esperado) {
		t.Errorf("Esperado %.2f, obteve %.2f", esperado, resultado)
	}
}

func TestCalculoSalario_CategoriaA(t *testing.T) {
	minutos := 240.0                              // 4 horas
	salarioBase := 4 * 3000.0                     // 12000
	esperado := salarioBase + (salarioBase * 0.5) // 18000
	resultado := calculoSalario(minutos, categoriaA)
	if !floatEquals(resultado, esperado) {
		t.Errorf("Esperado %.2f, obteve %.2f", esperado, resultado)
	}
}

func TestCalculoSalario_CategoriaInvalida(t *testing.T) {
	minutos := 60.0
	esperado := 0.0
	resultado := calculoSalario(minutos, "X")
	if resultado != esperado {
		t.Errorf("Esperado %.2f, obteve %.2f", esperado, resultado)
	}
}

func TestCalculoSalario_MinutosZero(t *testing.T) {
	esperado := 0.0
	resultadoA := calculoSalario(0, categoriaA)
	resultadoB := calculoSalario(0, categoriaB)
	resultadoC := calculoSalario(0, categoriaC)
	if resultadoA != esperado || resultadoB != esperado || resultadoC != esperado {
		t.Errorf("Esperado 0.00 para minutos zero, obteve %.2f, %.2f, %.2f", resultadoA, resultadoB, resultadoC)
	}
}
