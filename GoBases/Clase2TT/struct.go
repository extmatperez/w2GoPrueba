package main

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
)

type Fecha struct {
	Dia, Mes, Anio int
}

type Persona struct {
	Nombre    string `json:"nombrecito"`
	Apellido  string
	Direccion string `json:"direccion_persona"`
	Edad      int
	FechaNac  Fecha
}

// type entero int

func Concatenar(personita Persona) {
	fmt.Println(personita.Nombre + " " + personita.Apellido)
}

func (personita Persona) metodoConcatenar() {
	fmt.Println(personita.Nombre + " " + personita.Apellido)
}

type Circulo struct {
	Radio float64
	Area  float64
}

func (c *Circulo) setRadio(nuevoRadio float64) {
	c.Radio = nuevoRadio
}
func (c *Circulo) area() {
	c.setRadio(100)
	c.Area = math.Pi * c.Radio * c.Radio
	fmt.Println(c.Area)
}

func (c Circulo) getRadio() float64 {

	return c.Radio
}

func areaFuncion(c *Circulo) {
	fmt.Println("Dentro de areaFuncion")
	c.Area = math.Pi * c.Radio * c.Radio
	fmt.Println(c.Area)
}

func (c Circulo) perimetro() {
	fmt.Println(math.Pi * c.Radio)
}

func main() {
	// var numero1 int

	// numero1 = 25
	// var numero2 entero
	// numero2 = 32

	// fmt.Printf("\nnumero1: %v, %T, %d", numero1, numero1, unsafe.Sizeof(numero1))
	// fmt.Printf("\nnumero2: %v, %T, %d", numero2, numero2, unsafe.Sizeof(numero2))

	p1 := Persona{"Matias", "Perez", "Chilecito", 27, Fecha{17, 11, 1994}}
	fmt.Printf("\np1: %+v, %T", p1, p1)
	p2 := Persona{}
	p2.Nombre = "Juan"
	fmt.Printf("\np2: %+v, %T", p2, p2)
	p3 := Persona{
		Apellido:  "Soto",
		Nombre:    "Ramiro",
		Direccion: "Nose",
		Edad:      30,
	}

	fmt.Printf("\np3: %+v, %T", p3, p3)

	miJSON, err := json.Marshal(p3)
	fmt.Printf("\n")
	fmt.Println(miJSON)
	fmt.Println(string(miJSON))
	fmt.Println(err)

	tipo := reflect.TypeOf(5)

	fmt.Println(tipo.Name())
	fmt.Println(tipo.Kind())

	Concatenar(p3)
	p3.metodoConcatenar()

	circulo1 := Circulo{
		Radio: 2.5,
	}

	fmt.Printf("\n%+v\n", circulo1)
	circulo1.area()
	//circulo1.perimetro()

	//areaFuncion(&circulo1)

	//	fmt.Printf("\n%+v", circulo1)

	//circulo1.setRadio(15)

	//fmt.Printf("\n%+v", circulo1)
	// p3.Nombre = "No tengo idea"
	// fmt.Printf("\np3: %+v, %T", p3, p3)
	// fmt.Printf("\np3: El nombre es: %v", p3.Nombre)

	// var p4 Persona

	// fmt.Printf("\np4: %+v, %T", p4, p4)

	// fmt.Printf("\np4: El largo del nombre: %d", len(p4.Nombre))
	// p4.Nombre = "Matias"
	// fmt.Printf("\np4: El largo del nombre: %d", len(p4.Nombre))

	// p5 := Persona{
	// 	Apellido:  "Soto",
	// 	Nombre:    "Ramiro",
	// 	Direccion: "Nose",
	// 	Edad:      30,
	// 	FechaNac: Fecha{
	// 		Dia:  17,
	// 		Mes:  11,
	// 		Anio: 1994,
	// 	},
	// }

	// fmt.Printf("\np5: %+v, %T", p5, p5)
	// p5.FechaNac.Dia = 27
	// fmt.Printf("\np5: %+v, %T", p5, p5)

	// fmt.Printf("\nReflect\n")

	// tipo := reflect.TypeOf(p5)
	// valor := reflect.ValueOf(p5)

	// for i := 0; i < tipo.NumField(); i++ {
	// 	fmt.Println("Miembro:", tipo.Field(i))
	// 	fmt.Printf("\n%+v\n", tipo.Field(i))
	// 	fmt.Printf("\nValor %+v\n", valor.Field(i))
	// }
}
