package main

import "fmt"

func funcion() {
	fmt.Println("Cerre la base de datos")
}

func main() {

	num := 5
	func() {
		fmt.Println("Hola soy una función anonima 1")
	}()
	func() {
		fmt.Println("Hola soy una función anonima 1")
	}()
	defer func() {
		fmt.Println("Hola soy una función anonima 2")
	}()
	func() {
		fmt.Println("Hola soy una función anonima 3", num)
	}()
	num = 10
	defer func() {
		fmt.Println("Hola soy una función anonima 4", num)
	}()

	defer funcion()

	num = 15
	panic("Panic Generado")

	// defer func() {
	// 	fmt.Println("Hola soy una función anonima 2")
	// }()

}
