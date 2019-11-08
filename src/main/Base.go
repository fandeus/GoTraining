package main

import (
	//导包  导出报名
	"fmt"
	"math"
	"math/rand"
)

//01 函数的使用
func add(x int, y int) int {
	return x + y
}

//02 连续一个或者多个已命名形参类型相同时，除最后一个以外，其他都可以省略
func addTow(x, y int) int {
	return y + x
}

//03 函数的多返回值 一个函数可以有多个返回值
func swap(x, y string) (string, string) {
	return y, x
}

//04 go 的返回值可被命名 返回值的名称具有一定意义，它可以作文档使用
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//05 变量 var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后
//var 语句可以出现在包或者函数级别
//包级别的
var c, python, java bool

//06 变量的初始化
//变量的初始化可以包含初始值，每一个变量对应一个
//如果变量值已经存在 则可以省略类型 变量会从初始值中获得类型
var ii, jj int = 1, 3

func main() {
	fmt.Println("My favorite is,", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(12))
	fmt.Println(math.Pi)

	fmt.Println("3 + 4 =", add(3, 4))
	fmt.Println("4 + 5 =", addTow(4, 5))

	//同一个函数有个多个返回值时的接收方式
	a, b := swap("hello", "wold")
	e, _ := swap("hello", "test")
	fmt.Println(a, b)
	fmt.Println(e)
	fmt.Println(split(17))

	//var 函数级别
	var i int
	fmt.Println(i, c, python, java)

	//var 如果变量值已经存在 类型可以省略 变量会从初始值中获得
	var cpp, pythonTow, javaTow = true, false, "no !"
	fmt.Println(ii, jj, cpp, pythonTow, javaTow)
}
