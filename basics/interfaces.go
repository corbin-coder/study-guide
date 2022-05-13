package basics

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width  float64
	height float64
}

type cirle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c cirle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c cirle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...shape) float64 {
	var peri float64
	for _, s := range shapes {
		peri += s.perimeter()
	}
	return peri
}

func calculate() {
	r := rect{width: 10, height: 10}
	c := cirle{radius: 10}

	fmt.Println("Total Area: ", totalArea(r, c))
	fmt.Println("Total Perimeter: ", totalPerimeter(r, c))
}
