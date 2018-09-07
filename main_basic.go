package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is main")
	bl()
	b2()
}

//变量
func bl() {
	fmt.Println("＝＝变量举例＝＝")
	var str string
	var i int
	str = "str this is string"
	i = 10
	fmt.Println("str=", str)
	fmt.Println("i=", i)

	str1 := "str1"
	i1 := 1
	fmt.Println("str1=", str1)
	fmt.Println("i1=", i1)

	var (
		str3 string
		str4 string
		i3   int
		i4   int
	)
	str3 = "str3"
	str4 = "str4"
	i3 = 3
	i4 = 4
	//str3, str4, i3, i4 = "str3", "str4", 3, 4
	fmt.Println("var:", str3, str4, i3, i4)

	var (
		str5 = "str5"
		i5   = 5
	)
	fmt.Println("var:", str5, i5)

	var str6 string = "str6"
	fmt.Println(str6)

	var str66 = "str66"
	fmt.Println(str66)
}

//常量
func b2() {
	const str7 string = "this is const"
	fmt.Println(str7)

	const str8 = "str8"
	fmt.Println(str8)

	const (
		str9 = "str9"
		i9   = 9
	)
	fmt.Println(str9, i9)
}

//
