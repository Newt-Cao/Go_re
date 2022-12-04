package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	// 连接到Redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Redis connected error = ", err)
		return
	} else {
		fmt.Println("Redis connected successfully!")
	}
	// 连接服务调用完后需要及时关闭，防止资源占用
	defer conn.Close()
	// 想数据库添加一个数据
	_, err = conn.Do("set", "name", "vian")
	if err != nil {
		fmt.Println("Set string error = ", err)
	} else {
		fmt.Println("Set string successfully!")
	}
	// 从数据库获取一个元素，因为Do函数返回的是interface{}类型需要转换使用String方法，获取多个时，Strings
	data, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("Date gain error = ", err)
	} else {
		fmt.Println("Date gain successfully!\n", data)
	}
}
