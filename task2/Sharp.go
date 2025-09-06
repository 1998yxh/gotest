package main

import (
	"fmt"
	"math"
)

type Sharp interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	height float64
	length float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.length
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.length)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.radius
}

func main() {

	area := Rectangle{height: 5, length: 7}
	fmt.Printf("area: %f\n perimeter: %f\n", area.Area(), area.Perimeter())

	circle := Circle{radius: 3}
	fmt.Printf("circle: %f\n perimeter: %f\n", circle.Area(), circle.Perimeter())
}
