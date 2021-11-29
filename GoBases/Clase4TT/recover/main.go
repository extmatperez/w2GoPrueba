package main

import "fmt"

func isPair(num int) {

	defer func() {

		err := recover()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("No hay error")
		}
		panic("Hola soy un panic")
	}()

	if num%2 != 0 {
		panic("no es un numero par")
	}

	panic("es un numero par")

	//	fmt.Printf("El numero es par")
}

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("No hay error")
		}
	}()

	num := []int{1, 2, 3, 4}

	fmt.Println(num[5])
	fmt.Println(num[2])

	// isPair(5)
	// fmt.Printf("Terminó la ejecución")
}
