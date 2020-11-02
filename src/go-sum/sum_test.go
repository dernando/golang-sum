package main

import "testing"

func TestSum(t *testing.T) {

	var a, b = 5, 5
	var resultSum := sumInt(5,5)

	if resultSum != 10 {
		t.Errorf("Resultado esperado=10, resultado recebido %d", resultSum)
	}

}