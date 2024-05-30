package main

import (
	"fmt"
	"math"
)

/*
Structs & Methods

Circle
	- Radius (float64)
	- Area() => returns the area of the circle (Pi * Radius * Radius)

Rectangle
	- Length (float64)
	- Breadth (float64)
	- Area() => returns the area of the rectangle (Length * Breadth)
*/

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func PrintArea(x ?){
	fmt.Println("Area :", x.Area())
}

func main() {
	/*
		Create an instance of a circle with Radius=12
		Print the area of the circle
	*/
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)
	/*
		Create an instance of a rectangle with Length=10 & Breadth=12
		Print the area of the rectangle
	*/
	r := Rectangle{Length: 10, Breadth: 12}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)
}
