package main

import (
	"fmt"
)

func clcdisp(i int) {

	for a := 1; a <= i; a++ {
		for j := 1; j <= a; j++ {
			b := a * j
			fmt.Printf("%dx%d=%d\t\t\t\t", j, a, b)
		}
		fmt.Println()
	}

}

func main() {
	clcdisp(12)
}
