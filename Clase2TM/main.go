package main

import (
	"errors"
	"fmt"
	"unsafe"
)

const (
	Suma     = "+"
	Resta    = "-"
	Multi    = "*"
	Divi     = "/"
	Potencia = "^"
)

func sumarNumeros(numero1, numero2 float64) {
	fmt.Printf("La suma de %v y %v es: %v\n", numero1, numero2, numero1+numero2)
	fmt.Println("La suma de ", numero1, " y ", numero2, "es: ", numero1+numero2)
}

func sumarNumeros2(numero1, numero2 float64) float64 {
	return numero1 + numero2
}

func operaciones(operacion string, num1, num2 float64) float64 {

	switch operacion {
	case Suma:
		return num1 + num2
	case Resta:
		return num1 - num2
	case Multi:
		return num1 * num2
	case Divi:
		if num2 != 0 {
			return num1 / num2
		} else {
			return 0.0
		}

	}
	return -9999
}

func analizarNumero(numero int) {
	if numero == 0 {
		fmt.Println(numero, " es cero")
	} else if numero < 0 {
		fmt.Println(numero, " es negativo")
	} else {
		fmt.Println(numero, " es positivo")
	}
	numero = 28
}

func sumar(operacion string, variable ...float64) float64 {
	aux := 0.0
	if operacion == Multi {
		aux = 1.0
	}

	for _, value := range variable {
		aux = operaciones(operacion, aux, value)
	}
	return aux
}

func opSuma(num1, num2 float64) float64 {
	return num1 + num2
}
func opResta(num1, num2 float64) float64 {
	return num1 - num2
}
func opMulti(num1, num2 float64) float64 {
	return num1 * num2
}
func opPoten(num1, num2 float64) float64 {
	return num1 * num2
}
func opDivi(num1, num2 float64) float64 {
	if num2 == 0 {
		return 0
	}
	return num1 / num2
}

func orquestador(valores []float64, operador func(num1, num2 float64) float64) float64 {
	aux := 0.0
	for _, value := range valores {
		aux = operador(aux, value)
	}
	return aux
}

func operacionAritmetica(operador string, valores ...float64) float64 {
	switch operador {
	case Suma:
		return orquestador(valores, opSuma)
	case Resta:
		return orquestador(valores, opResta)
	}
	return 0.0
}

func multipleRetorno(num1, num2 float64) (float64, float64, float64, float64) {
	return opSuma(num1, num2), num1 - num2, opMulti(num1, num2), 0.0
}

func multipleRetorno2(num1, num2 float64) (suma, resta float64) {
	suma = num1 + num2
	resta = num1 - num2

	return
}

func permutar(num1, num2 float64) (float64, float64) {
	return num2, num1
}

func division(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0.0, errors.New("El denominador es cero!!!!")
	}
	return num1 / num2, nil
}

func main() {

	a, b, c, d := 1.0, 0.0, 6.0, 6.0

	// analizarNumero(a)
	// analizarNumero(a)
	// analizarNumero(b)
	// analizarNumero(c)
	// analizarNumero(d)
	sumarNumeros(float64(a), (float64(c)))
	sumarNumeros(float64(b), (float64(d)))

	fmt.Println(sumarNumeros2(float64(a), (float64(c))))

	fmt.Printf("El resultado de sumar %v y %v es %v\n", a, c, operaciones(Suma, a, c))
	fmt.Printf("El resultado de restar %v y %v es %v\n", a, c, operaciones(Resta, a, c))
	fmt.Printf("El resultado de multiplicar %v y %v es %v\n", a, c, operaciones("*", a, c))
	fmt.Printf("El resultado de dividir %v y %v es %v\n", a, c, operaciones("/", a, c))
	fmt.Printf("El resultado del modulo %v y %v es %v\n", a, c, operaciones("%", a, c))

	fmt.Printf("\nEl resultado de sumar 2,5,6,8,9,4 es %v", sumar(Suma, 2, 5, 6, 8, 9, 4))
	fmt.Printf("\nEl resultado de sumar 2,5,6,8 es %v", sumar(Suma, 2, 5, 6, 8))
	fmt.Printf("\nEl resultado de sumar 2,5,6,8,9,4 es %v", sumar(Multi, 2, 5, 6, 8, 9, 4))
	fmt.Printf("\nEl resultado de sumar 2,5,6,8 es %v", sumar(Multi, 2, 5, 6, 8))

	fmt.Printf("\n\nUsando funciones como parametros:")

	fmt.Printf("\nEl resultado de sumar 2,5,6,8,9,4 es %v", operacionAritmetica(Suma, 2, 5, 6, 8, 9, 4))
	fmt.Printf("\nEl resultado de sumar 2,5,6,8 es %v\n", operacionAritmetica(Resta, 2, 5, 6, 8))

	numero1, _, _, _ := multipleRetorno(5, 6)
	fmt.Println(numero1)
	fmt.Println(multipleRetorno(5, 6))

	fmt.Printf("\na: %v, c: %v", a, c)
	// aux := a
	// a = c
	// c = aux

	a, c = permutar(a, c)

	fmt.Printf("\na: %v, c: %v", a, c)

	fmt.Printf("\n\nUsando errores:\n")

	divi, err := division(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("La división es: ", divi)
	}

	type entero int64
	var numerito entero = 15
	var numerito2 int64 = 15
	fmt.Printf("\nnumerito: %v, %T, %v", numerito, numerito, unsafe.Sizeof(numerito))
	fmt.Printf("\nnumerito2: %v, %T, %v", numerito2, numerito2, unsafe.Sizeof(numerito2))
}
