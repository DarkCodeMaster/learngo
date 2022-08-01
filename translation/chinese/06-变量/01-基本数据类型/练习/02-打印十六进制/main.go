// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// 本练习是自选的

// ---------------------------------------------------------
// 练习: 打印十六进制
//
//  1. 使用十六进制打印0到9
//
//  2. 使用十六进制打印10到15
//
//  3. 使用十六进制打印17
//
//  4. 使用十六进制打印25
//
//  5. 使用十六进制打印50
//
//  6. 使用十六进制打印100
//
// 预期输出
//  0 1 2 3 4 5 6 7 8 9
//  10 11 12 13 14 15
//  17
//  25
//  50
//  100
//
// 提示
//  https://stackoverflow.com/questions/910309/how-to-turn-hexadecimal-into-decimal-using-brain
//
// https://simple.wikipedia.org/wiki/Hexadecimal_numeral_system
//
// ---------------------------------------------------------

func main() {
	// 示例:

	// 我将使用十六进制打印10
	fmt.Println(0xa)

	// 我将使用十六进制打印16
	// 0x10
	//   ^^-----  1 * 0 = 0
	//   |
	//   +------ 16 * 1 = 16
	//                  = 16
	fmt.Println(0x10)

	// 我将使用十六进制打印150
	// 0x96
	//   ^^-----  1 * 6 = 6
	//   |
	//   +------ 16 * 9 = 144
	//                  = 150
	fmt.Println(0x96)

	// 注释掉上面所有的代码, 然后
	// 在下面添加你自己的解答
}
