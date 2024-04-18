package main

import "testing"

func TestSuma(t *testing.T) {
	resultado := sumar(2, 3)
	esperado := 5
	if resultado != esperado {
		t.Errorf("incompatible no paso la prueba")
	}
}
