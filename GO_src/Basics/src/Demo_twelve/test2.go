package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Id    int
	Name  string
	Class string
}

func (this *Student) Subject() {
	fmt.Printf("%v Go to class\n", this.Name)
}

func MethodFunc(arg interface{}) {
	if reflect.TypeOf(arg).Kind() == reflect.Struct {
		numMethod := reflect.TypeOf(arg)

		for i := 0; i < numMethod.NumMethod(); i++ {
			meThod := numMethod.Method(i)

			fmt.Printf("%s:%v", meThod.Name, meThod.Type)

		}
	}
}

func main() {
	Vivian := Student{
		Id:    10086,
		Name:  "VivianChan",
		Class: "20软件",
	}

	MethodFunc(Vivian)

}
