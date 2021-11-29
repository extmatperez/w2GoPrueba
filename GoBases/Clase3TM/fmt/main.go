package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Persona struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func main() {

	num := 99

	fmt.Printf("%d %c %o %b %x %X %v %p", num, num, num, num, num, num, &num, &num)
	fmt.Printf("\n%-10d, %10d, %-15s, %15s", num, num, "Matias", "Matias")
	salida := fmt.Sprintf("%b", 123456)
	fmt.Println(salida)
	var err error
	err = os.Setenv("NAME", "gopher")
	fmt.Println(err)
	err = os.Setenv("NAME", "")
	//variable := os.Getenv("NAME")
	variable2, ok := os.LookupEnv("NAME")

	if ok {
		fmt.Printf("\nSe encontró la variable de entorno: %s\n", variable2)
	} else {
		fmt.Printf("\nNo se encontró la variable de entorno: %s\n", variable2)
	}
	// fmt.Println(variable)
	// fmt.Println(err)

	data, err := os.ReadFile("./archivo/myFile.txt")
	if err == nil {
		file := string(data)
		fmt.Println(file)
	} else {
		fmt.Println("El archivo no existe...")
	}

	// p1 := Persona{"Matias", "Perez"}
	// p2 := Persona{"Juan", "Perez"}
	p3 := Persona{"Mateo", "Perez"}

	// var lista []Persona

	// lista = append(lista, p1)
	// lista = append(lista, p2)

	// p1formateado, err := json.Marshal(lista)

	// err = os.WriteFile("./archivo/myFile2.txt", p1formateado, 0644)

	// if err != nil {
	// 	fmt.Println("No se pudo escribir")
	// }

	data, err = os.ReadFile("./archivo/myFile2.txt")

	var pListaLeida []Persona

	json.Unmarshal(data, &pListaLeida)

	fmt.Printf("\nEl nombre de la 2da persona en el archivo es: %s", pListaLeida[1].Nombre)
	fmt.Printf("\nEl apellido de la 2da persona en el archivo es: %s", pListaLeida[1].Apellido)
	fmt.Printf("\n%+v", pListaLeida)

	fmt.Fprintf(os.Stdout, "\nHola")

	pListaLeida = append(pListaLeida, p3)

	p1formateado, err := json.Marshal(pListaLeida)

	err = os.WriteFile("./archivo/myFile2.txt", p1formateado, 0644)

	if err != nil {
		fmt.Println("No se pudo escribir")
	}

	// if err == nil {
	// 	file := string(data)
	// 	fmt.Println(file)
	// } else {
	// 	fmt.Println("El archivo no existe...")
	// }
}
