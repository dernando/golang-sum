package main

import "testing"

func TestSum(t *testing.T) {

	resultSum := sumInt(5,5)

	if result != 10 {
		t.Errorf("Resultado esperado=10, resultado recebido %v", result)
	}

}