package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// 实例化全局Pool，函数体内实例化其他函数无法使用
var pool *redis.Pool

// 使用 ini{t 初始化
func init() {
	// redigo中的pool是一个结构体，所以创建结构体实例，下面是常规链接池配置
	pool = &redis.Pool{
		MaxIdle:     8,   // 池里存放的链接数，也叫最大空闲链接数
		MaxActive:   0,   // 和Redis连接的最大链接数，0表示无限制
		IdleTimeout: 300, // 空闲链接的最大空闲时间，时间结束后回收到链接池内
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		}, // 创建连接到指定数据库
	}
}

func main() {
	conn := pool.Get() // 从链接池中取一个链接
	defer conn.Close() // 函数结束后将链接放回链接池，还可以继续向连接池取链接
	defer pool.Close() // 关闭连接池，无法继续从链接池取链接
	_, err := conn.Do("hset", "animal", "name", "Elephant", "age", 20)
	if err == nil {
		fmt.Println("Successfully!")
	} // 放数据
	data, err := redis.Strings(conn.Do("hgetall", "animal"))
	if err == nil {
		fmt.Println("Successfully!", data)
	} // 读数据

	/*
		Successfully!
		Successfully! [name Elephant age 20]
	*/
}
