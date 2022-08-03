// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// 变量的命名规范

func main() {
	// 正确的声明
	var speed int
	var SpeeD int

	// 下划线也是允许的, 但是并不推荐
	var _speed int

	// Unicode字符也是允许的
	var 速度 int

	// 让编译器满意
	_, _, _, _ = speed, SpeeD, _speed, 速度
}
