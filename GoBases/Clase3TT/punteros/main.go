package main

import "fmt"

func main() {

	var p1 *int
	var num int

	var p2 = new(int)
	p2 = &num
	*p2 = 65
	fmt.Printf("p1Antes de inicializar: %v %v", p1, num)
	fmt.Printf("\np2Antes de inicializar: %v %v", p2, num)
	p1 = &num
	num = 5
	fmt.Printf("\nDespués de inicializar: %v %v", p1, *p1)
	fmt.Printf("\nDespués de inicializar: %v %v", &num, num)

	*p1 = 15
	fmt.Printf("\nDespués de inicializar: %v %v", p1, *p1)
	fmt.Printf("\nDespués de inicializar: %v %v", &num, num)
	fmt.Printf("\nDespués de inicializar: %v %v", p2, *p2)

	p3 := &p2

	p4 := &p3
	*(*(*p4)) = 20

	fmt.Printf("\nDespués de inicializar: %v %v", &num, num)

}
