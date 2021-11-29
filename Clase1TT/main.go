package main

import (
	"fmt"
)

func main() {

	// x, y := 21, 10

	// fmt.Printf("%v", x/y)
	// fmt.Printf("\n%v", x%y)

	// x++
	// fmt.Printf("\n%v", x)
	// x--
	// fmt.Printf("\n%v", x)

	// // x += 5

	// // x = x + 5

	// numero := 25

	// pNumero := &numero

	// fmt.Printf("\n%v, %T, %v", numero, numero, &numero)
	// fmt.Printf("\n%v, %T, %v, %v", pNumero, pNumero, &pNumero, *pNumero)

	// var a [5]int
	// var b []string
	// // a[0] = "Matias"
	// // a[1] = "Perez"
	// //a[2] = "Juan"

	// fmt.Printf("\n%v, %T", b, b)

	// var s = []bool{true, false, true}
	// fmt.Printf("\n%v", s)

	// slice := make([]int, 4)
	// fmt.Printf("\narray %v, %T", a, a)
	// fmt.Printf("\nslice: %v, %T", slice, slice)

	// primes := []int{3, 5, 7, 8, 9, 10}
	// fmt.Println(primes)
	// fmt.Println(primes[1:4])
	// fmt.Println(primes[1:])
	// fmt.Println(primes[:4])

	// slice[0] = 25

	// fmt.Printf("\nlen: %d, cap: %d", len(slice), cap(slice))

	// slice = append(slice, 2)

	// fmt.Printf("\nslice %v len: %d, cap: %d", slice, len(slice), cap(slice))

	// slice = append(slice, 3)
	// slice = append(slice, 3)
	// slice = append(slice, 3)
	// slice = append(slice, 3)
	// slice = append(slice, 3, 5, 6, 8)

	// fmt.Printf("\nslice %v len: %d, cap: %d", slice, len(slice), cap(slice))

	// slice = slice[1:]

	// fmt.Printf("\nslice %v len: %d, cap: %d", slice, len(slice), cap(slice))
	var myMap = map[string]string{}

	fmt.Printf("\nmyMap %v %T", myMap, myMap)
	fmt.Printf("\nmyMap['matias'] %v", myMap["matias"])
	myMap["matias"] = "27"
	myMap["juan"] = "25"
	myMap["rocio"] = "36"
	myMap["abel"] = "55"
	myMap["abel"] = "96"

	fmt.Printf("\nmyMap %v %T", myMap, myMap)
	fmt.Printf("\nmyMap['matias'] %v", myMap["matias"])
	// delete(myMap, "abel")
	// fmt.Printf("\nmyMap %v %T", myMap, myMap)
	// delete(myMap, "abel")
	// fmt.Printf("\nmyMap %v %T", myMap, myMap)

	for key, value := range myMap {
		fmt.Printf("\nclave: %v valor: %v", key, value)
	}

	frutas := []string{"banana", "manzana", "pera"}

	for i := 0; i < len(frutas); i++ {
		fmt.Printf("\n%d -> %v", i, frutas[i])
		frutas[i] = "naranjas"
	}

	for _, fruta := range frutas {
		fmt.Printf("\n%v", fruta)
		fruta = "nada"
	}

}
