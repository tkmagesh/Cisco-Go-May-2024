package main

import "fmt"

func main() {
	/* product => id, name, cost */
	// var product struct{}
	/*
		var product struct {
			Id   int
			Name string
			Cost float64
		}
		product.Id = 100
		product.Name = "Pen"
		product.Cost = 10
	*/

	/*
		var product struct {
			Id   int
			Name string
			Cost float64
		} = struct {
			Id   int
			Name string
			Cost float64
		}{
			100, "Pen", 10,
		}
	*/

	var product struct {
		Id   int
		Name string
		Cost float64
	} = struct {
		Id   int
		Name string
		Cost float64
	}{
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

func Format(p struct {
	Id   int
	Name string
	Cost float64
}) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(p *struct {
	Id   int
	Name string
	Cost float64
}, discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
