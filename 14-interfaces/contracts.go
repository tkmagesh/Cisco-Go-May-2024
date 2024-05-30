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

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Area() method not supported")
	}
}
*/

// runtime type check
/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }: // asserting if x is any object with Area() method
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Area() method not supported")
	}
}
*/

func PrintArea(x interface{ Area() float64 }) {
	fmt.Println("Area :", x.Area())
}

/*
Create a Perimeter method for both Circle & Rectangle
	Circle Perimeter = 2 * pi * radius
	Rectangle Perimeter = 2 * (Length + Breadth)

Create a "PrintPerimeter" function that will print the perimeter of the given object

In the main() function
	Use the PrintPerimeter function to print the perimeter of circle
	Use the PrintPerimeter function to print the perimeter of rectangle
*/

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

func PrintPerimeter(x interface{ Perimeter() float64 }) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// Printing Area() and Perimeter() together (contract / interface composition)
func PrintStats(x interface {
	interface{ Area() float64 }
	interface{ Perimeter() float64 }
}) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should be interface{ Perimeter() float64 }
}

func main() {
	/*
		Create an instance of a circle with Radius=12
		Print the area of the circle
	*/
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintStats(c)

	/*
		Create an instance of a rectangle with Length=10 & Breadth=12
		Print the area of the rectangle
	*/
	r := Rectangle{Length: 10, Breadth: 12}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintStats(r)
}
