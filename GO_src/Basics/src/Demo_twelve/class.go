/* 类的表示与封装 */

package main

import "fmt"

// 方法名、类名、包名、类型名，首字母大写表示公共访问和包内部访问，小写表示包内部访问
type Hero struct {
	Name  string
	Ad    int
	Level int //与上方类型一致时可省略数据类型
}

// 不加 * 号的 this 表示的是值拷贝，无法修改对象中的信息。
func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name, "Ad = ", this.Ad, "Level = ", this.Level)
}

//(this 结构体名) 表示该函数方法绑定到结构体，类方法。
func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	// 定义一个Hero结构体的对象
	hero := Hero{Name: "张三", Ad: 150, Level: 5}
	hero2 := Hero{Name: "李四", Ad: 200, Level: 10}

	hero.Show()
	hero.GetName()

	hero2.Show()
	hero2.SetName("王五")
	hero2.Show()

}
