package main

import "fmt"

func main() {
	fmt.Println("Hola mundo!!!")

	var edad int

	edad = 27

	fmt.Println(edad)
	fmt.Printf("La edad es: %v, %d, %T", edad, edad, edad)

	horas := '@'

	fmt.Printf("\nLas \" horas son: %v, %c, %T", horas, horas, horas)
	fmt.Printf("\nLas \" horas son: %v, %c, %T", horas, horas, horas)
	/*


		horas = 65

		fmt.Printf("\nLas \" horas son: %c, %.2v, %T", horas, horas, horas)
		horas = 1235
		fmt.Printf("\nLas \" horas son: %08.2f, %v, %T", horas, horas, horas)
		horas = 12365
		fmt.Printf("\nLas \" horas son: %8.2f, %v, %T\n", horas, horas, horas)

		fmt.Println("Las horas son:", horas, " y las horas son", horas)
	*/
}
