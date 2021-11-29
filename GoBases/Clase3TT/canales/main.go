package main

import (
	"fmt"
	"time"
)

func proceso(i int, c chan int) {
	fmt.Println(i, "-inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-termina")
	c <- i
}

func main() {
	c := make(chan int)
	suma := 0
	ini := time.Now()
	for i := 0; i < 10; i++ {
		go proceso(i, c)
		// <-c
	}
	for i := 0; i < 10; i++ {
		variable := <-c
		suma += variable
		fmt.Println(variable)
	}

	fin := time.Now()

	tiempo := fin.Sub(ini)

	fmt.Println("El tiempo paralelo demorado es de: ", tiempo.Seconds())
	fmt.Println("La suma es:  ", suma)
}
