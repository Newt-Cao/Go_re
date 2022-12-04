package factory

// 当factory里结构体Students为大写时，直接调用
type Students struct {
	Name   string
	Age    int
	Scores float64
}

// 当factory里结构体students为小写时，需要调用该包的大写函数间接调用小写结构体
type students struct {
	Name   string
	age    int
	Scores float64
}

// 返回一个students的指针类型所以间接调用了该结构体，函数体也是执行了该结构体的实例化对象，最后返回该结构体的指针：var stu *students = &students{...}
func StudentsInterface(name string, age int, scores float64) *students {
	return &students{
		Name:   name,
		age:    age,
		Scores: scores,
	}
}

// 当factory里结构体students为小写，结构体里字段也为小写时，需要绑定一个结构体大写函数间接调用小写结构体和结构体里小写字段
func (stu *students) AgeField() int {
	return stu.age
}
