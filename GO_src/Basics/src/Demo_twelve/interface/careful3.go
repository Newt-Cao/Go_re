package main

type One interface {
	test1()
}

type Two interface {
	test2()
}

type Three interface {
	One
	Two
	test3()
}

type Four struct {
}

// 必须实现 Three 所有方法
func (F *Four) test1() {
}

func (F *Four) test2() {
}

func (F *Four) test3() {
}

func main() {
	var four Four
	var three Three = &four
	three.test1()
	three.test2()
	three.test3()
}
