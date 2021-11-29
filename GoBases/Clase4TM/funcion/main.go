package main

import (
	"errors"
	"fmt"
)

func sumar(num1, num2 int) int {
	return num1 + num2
}

func restar(num1, num2 int) int {
	return num1 - num2
}

func operacion(opcion int) (func(a, b int) int, error) {
	if opcion == 1 {
		return sumar, nil
	} else if opcion == 2 {
		return restar, nil
	}
	return nil, errors.New("La opcion no es correcta")
}

func main() {
	funcion, err := operacion(8)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El resultado de usar la funciones: %d", funcion(5, 2))
	}

}
