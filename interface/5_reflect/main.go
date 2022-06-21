package main

import (
	"fmt"
	"math"
	"reflect"
)

// interfaces are named collections of method signatures
type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// interfaces are implemented implicitly
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func myReflect(g geometry) {
	fmt.Println(reflect.TypeOf(g))
	fmt.Println(reflect.ValueOf(g))
}

func main() {

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	myReflect(r)
	myReflect(c)
}
