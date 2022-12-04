package main

func main() {
	age := "abc"
	println(&age) //0xc000063f60
	age = "cbdeeeeeeeeee"
	println(&age) //0xc000063f60
}
