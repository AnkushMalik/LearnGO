package main

import (
	"fmt"
	"math"
)

// Interface Define
type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rect struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rect) area() float64 {
	return r.width * r.height
}
func getArea(s Shape) float64 {
	return s.area()
}
func main() {
	circle := Circle{7}
	rectangle := Rect{5, 7}
	fmt.Printf("Circle Area : %f\n", getArea(circle))
	fmt.Printf("Rectangle Area : %f\n", getArea(rectangle))

}
