package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	// 连接数据库
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Redis connected successfully!")
	}
	// 及时关闭
	defer conn.Close()
	// 一次写入多个hash数据
	_, err = conn.Do("hmset", "hero", "name", "simple", "age", 20, "EX", 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Date writed successfully!")
	}
	// 读取数据
	data, err := redis.Strings(conn.Do("hmget", "hero", "name", "age"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Date readed successfully!")
	}
	fmt.Println(data)
	// 遍历数据
	for i, v := range data {
		fmt.Println(i, ":", v)
	}
}
