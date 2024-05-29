package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	/* product => id, name, cost */
	// var product struct{}
	/*
		var product Product
		product.Id = 100
		product.Name = "Pen"
		product.Cost = 10
	*/

	/*
		var product Product  = Product {
			100, "Pen", 10,
		}
	*/

	// structs as "values"
	/*
		var p1 = Product{Id: 999, Name: "Stapler", Cost: 199}
		var p2 = Product{Id: 999, Name: "Stapler", Cost: 199}
		fmt.Println(p1 == p2)
	*/

	var p1 = Product{Id: 999, Name: "Stapler", Cost: 199}
	var p2 = p1 //creating a copy coz structs as "values"
	p2.Cost = 399
	fmt.Println("p1 :", p1)
	fmt.Println("p2 :", p2)

	var product Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	// fmt.Println(product)
	fmt.Printf("%+v\n", product)
	fmt.Println(Format(product))
	fmt.Println("After applying 10% discount")
	ApplyDiscount(&product, 10)
	fmt.Println(Format(product))
}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(p *Product, discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
