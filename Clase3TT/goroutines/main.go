package main

import (
	"fmt"
	"runtime"
	"time"
)

func proceso(i int) {
	fmt.Println(i, "-inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-termina")
}

func proceso2(i int) {
	fmt.Println(i, "-inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-termina")
}

func main() {
	// ini := time.Now()
	// proceso(1)
	// proceso(2)
	// proceso(3)
	// proceso(4)
	// fin := time.Now()
	// tiempo := fin.Sub(ini)

	// fmt.Println("El tiempo secuencial demorado es de: ", tiempo.Seconds())

	ini := time.Now()

	fmt.Println("El numero de CPu es: ", runtime.NumCPU())
	for i := 0; i < 12; i++ {
		go proceso(i)
	}
	time.Sleep(2000 * time.Millisecond)
	fin := time.Now()
	tiempo := fin.Sub(ini)

	fmt.Println("El tiempo paralelo demorado es de: ", tiempo.Seconds())
}
