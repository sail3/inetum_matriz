package main

import (
	"reflect"
	"testing"
)

// TestRotarMatriz2x2: ejecuta una rotacion de matriz de 2x2
func TestRotarMatriz2x2(t *testing.T) {
	input, output := [][]int{
		{1, 2},
		{3, 4},
	}, [][]int{
		{2, 4},
		{1, 3},
	}
	result := RotarMatriz(input)
	if reflect.DeepEqual(output, result) {
		t.Log("El test de Rotacion de matriz 2x2 ha pasado con exito")
	} else {
		t.Error("La rotacion de la matriz 2x2 ha tenido errores")
		t.Fail()
	}
}

// TestRotarMatriz3x3: Ejecuta un cambio de matriz de 3x3
func TestRotarMatriz3x3(t *testing.T) {
	input, output := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}, [][]int{
		{3, 6, 9},
		{2, 5, 8},
		{1, 4, 7},
	}
	result := RotarMatriz(input)
	if reflect.DeepEqual(output, result) {
		t.Log("El test de Rotacion de matriz 3x3 ha pasado con exito")
	} else {
		t.Error("La rotacion de la matriz 3x3 ha tenido errores")
		t.Fail()
	}
}
