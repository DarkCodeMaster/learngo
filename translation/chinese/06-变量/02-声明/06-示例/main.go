// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	// 命名是大小写敏感的:
	// MyAge, myAge, and MYAGE 都是不同的变量

	// 使用案例:
	// 什么时候使用并行声明?
	//
	// 不太好:
	// var myAge int
	// var yourAge int
	//
	// 还可以:
	// var (
	// 	myAge int
	// 	yourAge int
	// )
	//
	// 比较好:
	var myAge, yourAge int
	fmt.Println(myAge, yourAge)

	var temperature float64
	fmt.Println(temperature)

	var success bool
	fmt.Println(success)

	var language string
	fmt.Println(language)
}
