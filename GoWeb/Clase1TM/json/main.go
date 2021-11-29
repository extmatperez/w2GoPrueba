package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre              string
	Apellido            string `json:"apellidito"`
	FechaNacimiento     string
	Direccion_domicilio string `json:"direccion_domicilio"`
}

func main() {

	fmt.Println(5)
	fmt.Println(5)

	// var s []int

	// s = append(s, 1, 2, 3, 4, 5)

	per := persona{"Matias", "Perez", "17/11/1994", "Chilecito"}
	var per2 persona
	salida, _ := json.Marshal(per)

	fmt.Println(salida)

	cadena := `{"nombre":"25","Apellido":25,"FechaNacimiento":"17/11/1994","direccion_domicilio":"Chilecito"}`

	fmt.Println(string(salida))

	err := json.Unmarshal([]byte(cadena), &per2)

	if err != nil {
		fmt.Println("Error!!!")
	} else {
		fmt.Println(per2)
	}

	// p = nil
	// salida, _ = json.Marshal(p)

	// fmt.Println(salida)
	// fmt.Println(string(salida))

}
