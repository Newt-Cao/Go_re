package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	b := "Vian"

	aReflect := reflect.ValueOf(a).Int()
	fmt.Printf("aReflect type = %T\naReflect value = %v\n", aReflect, aReflect)

	c := a + int(aReflect)
	fmt.Printf("c type = %T\nc value = %v\n", c, c)

	bReflect := reflect.ValueOf(b).String()
	fmt.Printf("bReflect type = %T\nbReflect value = %s\n", bReflect, bReflect)
}

/*
aReflect type = int64
aReflect value = 10
c type = int
c value = 20
bReflect type = string
bReflect value = Vian
*/
