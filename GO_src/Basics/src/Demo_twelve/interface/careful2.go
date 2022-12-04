package main

import "fmt"

type Oneinterface interface {
	Sum(n1, n2 integer) integer
}

type Twointerface interface {
	Sum(n1, n2 integer) integer
}

type integer int

func (num *integer) Sum(n1, n2 integer) integer {
	return n1 + n2
}

func main() {
	var Int integer
	var oneinterface Oneinterface = &Int
	fmt.Println(oneinterface.Sum(1, 2)) // 3

	var twointerface Twointerface = &Int
	fmt.Println(twointerface.Sum(3, 4)) // 7
}
