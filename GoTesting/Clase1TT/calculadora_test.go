package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	//Arrange
	num1 := 3
	num2 := 1

	//Act
	resultado := Dividir(num1, num2)

	//Assert
	assert.NotNil(t, resultado)
}

func TestSumar(t *testing.T) {
	//Arrange
	num1 := 3
	num2 := 10
	resultadoEsperado := 13

	//Act
	resultado := Sumar(num1, num2)

	//Assert
	if resultado != resultadoEsperado {
		t.Errorf("La funcion sumar obtuvo %v pero el resultado esperado era %v", resultado, resultadoEsperado)
	}
}

func TestRestar(t *testing.T) {
	//Arrange
	num1 := 3
	num2 := 10
	resultadoEsperado := -7

	//Act
	resultado := Restar(num1, num2)

	//Assert
	if resultado != resultadoEsperado {
		t.Errorf("La funcion restar obtuvo %v pero el resultado esperado era %v", resultado, resultadoEsperado)
	}
}

func TestSumar2(t *testing.T) {
	//Arrange
	num1 := 3
	num2 := 10
	resultadoEsperado := 13

	//Act
	resultado := Sumar(num1, num2)

	//Assert
	assert.Equal(t, resultadoEsperado, resultado, "Los valores no son iguales")
}
