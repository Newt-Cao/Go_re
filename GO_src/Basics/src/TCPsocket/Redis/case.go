package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Monster struct {
	name  string
	age   int
	skill string
}

func main() {
	var (
		name  string
		skill string
		age   int
	)
	monster := &Monster{}
	fmt.Print("Monster name：")
	fmt.Scan(&name)
	fmt.Print("Monster age：")
	fmt.Scan(&age)
	fmt.Print("Monster skill：")
	fmt.Scan(&skill)
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
	}
	_, err = conn.Do("hmset", "monster", monster.name, name, monster.age, age, monster.skill, skill)
	if err == nil {
		fmt.Println("successfully!")
	}
	data, err := redis.String(conn.Do("hgetall", "monster"))
	if err == nil {
		fmt.Println("Successfully!")
		fmt.Println(data)
	}
}
