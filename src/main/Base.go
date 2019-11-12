package main

import (
	//导包  导出报名
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

//01. 函数的使用
func add(x int, y int) int {
	return x + y
}

//02. 连续一个或者多个已命名形参类型相同时，除最后一个以外，其他都可以省略
func addTow(x, y int) int {
	return y + x
}

//03. 函数的多返回值 一个函数可以有多个返回值
func swap(x, y string) (string, string) {
	return y, x
}

//04. go 的返回值可被命名 返回值的名称具有一定意义，它可以作文档使用
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//05. 变量 var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后
//var 语句可以出现在包或者函数级别
//包级别的
var c, python, java bool

//06. 变量的初始化
//变量的初始化可以包含初始值，每一个变量对应一个
//如果变量值已经存在 则可以省略类型 变量会从初始值中获得类型
var ii, jj int = 1, 3

//07. 短变量的声明
//在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明
//函数外的每句语句都必须以关键开始 (var, func 等等)，因此 := 结构不能在函数外使用
var ijk = 23

// ijks := 23

//08. go 的基本类型
//表示 与或的 boole
//字符串 string
//int int8 int16 int32 int64
//uint uint8 uint16 uint32 uint64 uintptr
//byte //unit 的别名
//rune int32的别名 表示一个unicode点
//float32 float64
//complex64	complex128
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//09. 零值
//数值类型 0
//boole 类型为false
//string 为"" (空字符串)
var i1 int
var f float64
var bb bool
var s string

//10. 类型转换
//表达式 T(v) 将v转换成类型T
/**
  关于数值的转换
  var i int  = 42
  var f float64 = float64(1)   //将int 转换成 float64
  var u uint = uint(f)

  或者更简单的形式
  i := 42
  f := float64(i)
  u := uint(f)

*/

//11. 类型推导
//在声明一个对象 不指定其类型时
//即使用不带类型的 := 语法 或 var = 表达式 时，变量的类型由右边的值推导得出
//当右边的值声明了类型时，新变量的类型与其相同
//var i11 int
//j11 := i11  //j11 也是一个 int

//不过当右边包含未指明类型的数值常量时，新的变量就可能是int，float64 或者complex128 了这取决于
//常量的精度
/**
i := 42      //int
f := 3.142   //float64
g := 0.86 + 0.5i   //complex128
*/

//12. 常量
//常量的声明与变量类似，只不过使用 const 关键字
//常量可以是字符、字符串、 布尔值或数值
//常量不能使用 := 语法声明
const Pi = 3.14

//13.数值常量
/**
 数值常量是高精度的值
一个未指定类型的常量由上下文来决定其类型
int 类型最大可以存储一个64位的整数，有时会更小
int 可以存放最大64位的整数，根据平台的不同有时会更小
*/

const (
	//将 1 左移动100位来创造一个非常大的数值
	//即这个数的二进制是 1 后面跟着100 个 0
	Big = 1 << 100
	//再往右移 99 位，即 Small = 1 << 1 , 或者 Small = 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}


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

	// := 简单赋值不能在函数外使用
	tempVar := 34
	fmt.Println(ijk, tempVar)

	//基本类型
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// 零值
	fmt.Printf("Type: %T , int= %v\n", i1, i1)
	fmt.Printf("Type: %T, float64 = %v\n", f, f)
	fmt.Printf("Type: %T, bool = %v\n", bb, bb)
	fmt.Printf("Type: %T, string = %q\n", s, s)
	fmt.Printf("%v %v %v %q\n", i1, f, bb, s)

	//类型转换
	var x10, y10 int = 12, 13
	var f10 float64 = float64(x10*x10) + float64(y10*y10)
	fmt.Printf("Type: %T, Value: %v\n", f10, f10)
	var z10 uint = uint(f)
	fmt.Printf("Type: %T, Value: %v\n", z10, z10)

	//类型推导
	v := 4.2 + 0.5i // 修改这里推导类型
	fmt.Printf("v is of type %T\n", v)

	//常量
	const World = "世界"
	fmt.Println("hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)

	//数值常量
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
