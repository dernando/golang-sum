package main

import "testing"

func TestSum(t *testing.T) {

	var a, b = 5, 5
	var resultSum = sumInt(a,b)
	var expected = 10

	if resultSum != expected {
		t.Errorf("Resultado esperado=%d, resultado recebido %d", expected, resultSum)
	}

}