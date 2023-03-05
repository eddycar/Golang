package main

import (
	"fmt"
	"math"
)

const (
	rectType   = "RECT"
	circleType = "CIRCLE"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
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

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func details(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

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

	r := newGeometry(rectType, 2, 3)
	details(r)
	c := newGeometry(circleType, 2)
	details(c)

}
