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

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func main() {
	var product Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	fmt.Printf("%+v\n", product)
	// "Format" as a function
	// fmt.Println(Format(product))

	// "Format" as a method
	fmt.Println(product.Format())

	fmt.Println("After applying 10% discount")

	// "ApplyDiscount" & "Format" as a functions
	// ApplyDiscount(&product, 10)
	// fmt.Println(Format(product))

	// "ApplyDiscount" & "Format" as a methods
	product.ApplyDiscount(10)
	fmt.Println(product.Format())

}
