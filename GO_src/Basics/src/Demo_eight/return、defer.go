/* return与defer的执行顺序 */

package main

import "fmt"

// return 先执行 ，defer 后执行

func deferfunc() int {
	fmt.Println("deferfunc run")
	return 0
}

func returnfunc() int {
	fmt.Println("returnfunc run")
	return 0
}

func Dcision() int {

	defer deferfunc()

	return returnfunc()
}

func main() {

	Dcision()

} // returnfunc run  deferfunc run
