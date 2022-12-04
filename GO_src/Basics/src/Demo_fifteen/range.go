package main

import "fmt"

func main() {
	c := make(chan int, 5)

	go func() {
		for i := 0; i < cap(c); i++ {
			c <- i
		}

		close(c)
	}()

	// 使用range迭代不断操作的channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main finished.")
}
