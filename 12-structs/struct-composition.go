package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

type Dummy struct {
	Id int
}

type PerishableProduct struct {
	Product
	// Dummy
	Expiry string
}

// factory function
func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	/*
		grapes := PerishableProduct{
			Product: Product{
				Id:   100,
				Name: "Grapes",
				Cost: 99,
			},
			Expiry: "5 Days",
		}
	*/

	// Use the factory function
	grapes := NewPerishableProduct(100, "Grapes", 99, "5 Days")
	fmt.Printf("%+v\n", grapes)
	// fmt.Println(grapes.Product.Id)
	fmt.Println(grapes.Id)
}
