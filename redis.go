package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	defer c.Close()

	_, err = c.Do("SET", "mykey", "wwel123")

	if err != nil {
		fmt.Println("SET to redis error", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("redis get failed: %v \n", username)
	}

	//查看redis已有数据量
	size, err := c.Do("DBSIZE")
	fmt.Printf("size is %d \n", size)
	//执行set命令，写入数据
	_, err = c.Do("set", "name", "yuanye")
	if err != nil {
		checkErr(err)
	}
	//取数据
	name, err := redis.String(c.Do("get", "name"))
	if err != nil {
		checkErr(err)
	} else {
		fmt.Println(name)
	}
	//删除数据
	_, err = c.Do("del", "name")
	if err != nil {
		checkErr(err)
	}
	//检查name是否存在
	has, err := redis.Bool(c.Do("exists", "name"))
	if err != nil {
		fmt.Println("name is", err)
	} else {
		fmt.Println(has)
	}
	//设置redis过期时间3s
	_, err = c.Do("set", "myName", "hehe")
	if err != nil {
		checkErr(err)
	}
	myName, err := redis.String(c.Do("get", "myName"))
	if err != nil {
		checkErr(err)
	} else {
		fmt.Println("myName : ", myName)
	}
	//5s后取数据
	time.Sleep(time.Second * 5)
	myName, err = redis.String(c.Do("get", "myName"))
	if err != nil {
		fmt.Println("After 5s ", err)
	} else {
		fmt.Println("After 5s myName : ", myName)
	}
}

func checkErr(err error) {
	fmt.Println(err)
}
