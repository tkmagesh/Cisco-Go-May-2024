package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func main() {
	var x interface{}
	x = 100
	fmt.Println("x :", x)

	x = "Aliqua labore et enim consectetur voluptate deserunt consequat."
	fmt.Println("x :", x)

	x = true
	fmt.Println("x :", x)

	x = 99.86
	fmt.Println("x :", x)

	/* unsafe */
	// x = 100
	// x = "Mollit non sint deserunt eiusmod sit pariatur ullamco do sit labore et velit magna."
	// y := 200 + x.(int)
	// fmt.Println(y)

	/* safe */
	x = 100
	// x = "Mollit non sint deserunt eiusmod sit pariatur ullamco do sit labore et velit magna."
	if val, ok := x.(int); ok {
		y := 200 + val
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	/* using "type switch" */
	// x = 100
	// x = "Exercitation cupidatat aliquip minim do reprehenderit nostrud est eu enim elit labore excepteur."
	// x = true
	// x = 23.45
	x = Product{Id: 100, Name: "Pen", Cost: 10}
	// x = 4 + 5i

	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float64, 99.99% of x =", val*0.9999)
	case Product:
		fmt.Println("x is a product, ", val.Format())
	default:
		fmt.Println("x is of unknown type")
	}

}
