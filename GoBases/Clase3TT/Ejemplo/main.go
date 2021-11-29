package main

import "fmt"

type humano struct {
	Nombre   string
	Apellido string
}

func main() {
	personita := &humano{"Pepe", "Morales"}
	fmt.Println((*(&(*personita))).Nombre)
}
