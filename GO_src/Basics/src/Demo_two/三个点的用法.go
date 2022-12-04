package main

import "fmt"

func Sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}

func main() {
	/*
		1.用在函数参数中，若函数的最后一个参数是…T类型，这个参数可传入任意个T类型的参数，在函数中…T的类型为[]T.
	*/
	func(num ...int) {
		res := 0
		for _, v := range num {
			res += v
		}
		fmt.Println(res) // 15
	}(1, 2, 3, 4, 5)

	/*
		2.用在解序列,可以传入一个slice，然后用…解开它，注意在这里没有新的slice被创造。传入切片。
	*/
	primes := []int{2, 3, 5, 7}
	fmt.Println(Sum(primes...)) // 17

	/*
		3.append()函数传入切片时。
	*/
	var a = []int{1, 2}
	var b []int
	b = append(b, a...)
	fmt.Println(b) // [1 2]

	/*
		4.数组的定义。
	*/
	var array = [...]int{1, 2, 3}
	fmt.Println(array) // [1,2,3]
}
