package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 这里定义的变量不能使用:= 这里作用域是包内部的 go语言没有全局变量
var (
	aa = 3
	ss = "kkk"
	bb = true
)

// 也可以这样定义
var cc, dd, ee = 3, "kkk", true

func variableZeroValue() {
	var a int
	var b int
	fmt.Printf("%d,%q\n", a, b)
}

// hello world
// 0,'\x00' 使用%q打印出空串

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
} // 3 4 abc

func variableTypeDeduction() {
	// 在使用自动类型推断的时候可以将不同类型的变量放到一行上进行赋值
	var a, b, c, d = 3, 4, true, "def"
	fmt.Println(a, b, c, d)

} // go 语言中的类型自动推断(type deduction)
// 3 4 true def

func variableShorter() {
	// 这里定义的变量作用域是函数内
	a, b, c, d := 3, 4, true, "def"
	// 第一次定义变量可以使用 :=
	b = 5
	fmt.Println(a, b, c, d)
}

// 复数
func euler() {
	// 验证欧拉公式
	//c := cmplx.Pow(math.E, 1i*math.Pi) + 1
	c := cmplx.Exp(1i*math.Pi) + 1
	fmt.Println(c)
} // (0+1.2246467991473515e-16i)
// 实部是0 虚部接近0 原因是浮点数的不精确造成的 complex64 实部是float32 虚部也是float32

func main() {
	fmt.Print("hello world\n")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)
	euler()
}
