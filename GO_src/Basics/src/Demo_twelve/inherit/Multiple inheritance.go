package main

import "fmt"

type Shop struct {
	Goods string
	Price float64
}

type Books struct {
	Kind   string
	Weight float64
}

type Customer struct {
	Shop
	Books
	int
}

func main() {
	var customer Customer
	customer.int = 10
	fmt.Println(customer.int) // 10
}
