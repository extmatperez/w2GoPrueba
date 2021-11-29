package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
	nombre() string
}

type rect struct {
	width, height float64
}

type cuadrado struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (r rect) nombre() string {
	return "Rectangulo"
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2)
}

func (c circle) perim() float64 {
	return math.Pi * 2 * c.radius
}

func (c circle) nombre() string {
	return "Circulo"
}

func (c circle) nombre2() string {
	return "Circulo"
}

func details(g geometry) {
	fmt.Printf("%T", g)
	// fmt.Println(g.nombre())
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

const (
	rectType   = "RECT"
	circleType = "CIRCLE"
)

func newGeometry(geoType string, values ...float64) geometry {
	switch geoType {
	case rectType:
		return rect{width: values[0], height: values[1]}
	case circleType:
		return circle{radius: values[0]}
	}
	return nil
}

func main() {
	r := rect{width: 3, height: 4}
	//cuadr := cuadrado{width: 3, height: 4}
	c := circle{radius: 5}

	fmt.Printf("\nArea circulo: %.2f\n", c.area())

	details(r)
	details(c)
	//details(cuadr)

}
