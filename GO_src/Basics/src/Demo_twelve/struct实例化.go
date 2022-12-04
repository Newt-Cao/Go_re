package main

import "fmt"

func main() {
	type Person struct {
		Name string
		age  int
	}
	// 第一种
	var oneUser Person    // 会有默认值
	oneUser.Name = "Vian" // 也可以赋值
	fmt.Println(oneUser)  // {Vian 0}

	// 第二种
	twoUser := Person{"Jack", 20} // 推导后直接初始化
	threeUser := Person{}         // 也可以单独赋值
	threeUser.Name = "s1mple"
	threeUser.age = 35
	fmt.Println(twoUser)   // {Jack 20}
	fmt.Println(threeUser) // {s1mple 35}

	// 第三种
	fourUser := new(Person)   // 推导一个指向Person的指针，相当于 *Person 的指针类型
	(*fourUser).Name = "Niko" // 正常赋值，因为是指针，所需需要取指针的值*fourUser
	fourUser.age = 30         // 开发者优化后，不需要取值，系统底层自动实现，可直接赋值
	fmt.Println(*fourUser)    // {Niko 30}

	// 第四种(形式类似第三种，使用类似第二种)
	fiveUser := &Person{"Mary", 18} // 同样可以直接赋值，等价于 var fiveUser *Person = &Person{}
	fiveUser.age = 20               // 也可以分开赋值，也同样因为优化，所以不需要取值*
	fmt.Println(*fiveUser)          // {Mary 20}

	// 声明一点：new()和指针类型的区别是，new()已经开辟好内存空间，默认值为字段默认值，而指针类型没有开辟，默认为nil，需要传入&指针
}
