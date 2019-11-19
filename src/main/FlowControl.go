package main

import (
	"fmt"
	"math"
)

/**
第二章：流程控制
*/
//01. Go 只有一种循环结构： for 循环
//基本的 for 循环由部分组成，他们用分号隔开
// 1.初始化语句：在第一次迭代前执行
// 2.条件表达式：在每次迭代前求值
// 3.后置语句：在每次迭代结尾执行
// 初始化语句通常为一句短变量声明，改变量声明仅在	for 语句的作用域中可见
// 一旦条件表达式的布尔值为false，循环就终止
// 注意：Go 的循环和 C，java，javaScript 之类的语言不同，Go的for语句后面的构成部分没有小括号，大括号{} 是必须的

//02. Go 的 for 后续
//Go 的初始化语句和后置语句是可选的

//03. for 是 Go 中的 while
//此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for

//04. 无限循环
//如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑

//05. if
//GO的if语句和for循环类似，表达式外无需小括号（），然而大括号是必须的
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//06. if 的简短语句
//痛for一样，if语句可以在条件表达式前执行一个简单的语句
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

//07. if 和 else
func powIfElse(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	//这里就不能再使用v了
	return lim
}

func main() {

	//for 循环
	sum := 10
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//for 循环后续
	sum2 := 1
	for ; sum2 < 1000; {
		sum2 += sum2
	}
	fmt.Println(sum2)

	//go 中的 while
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)

	//infinite loop 无限循环
	/*
		for {
			fmt.Println("无限循环")
		}
	*/

	//if
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(2, 3, 10),
		pow(2, 3, 20),
	)

	//if 和 else
	fmt.Println(
		powIfElse(3, 2, 10),
		powIfElse(3, 2, 20),
	)

	// https://tour.go-zh.org/flowcontrol/8
}
