// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// 练习: 让我们运行这个函数自己看看零值

func main() {
	var speed int    // 数值类型
	var heat float64 // 数值类型
	var off bool
	var brand string

	fmt.Println(speed)
	fmt.Println(heat)
	fmt.Println(off)

	// 我使用printf打印过空字符串
	// 练习: 使用Println看看会发生啥
	fmt.Printf("%q\n", brand)
}
