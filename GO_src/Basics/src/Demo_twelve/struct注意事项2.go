package main

type A struct {
	B *A
}

func main() {
	a := &A{}
	b := &A{}
	a.B = b
}
