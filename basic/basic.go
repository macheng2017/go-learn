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
	fmt.Printf("%.3f", c)
} // (0.000+0.000i)
// 实部是0 虚部接近0 原因是浮点数的不精确造成的 complex64 实部是float32 虚部也是float32
// 在go中只有强制类型转换(显式转换)
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Printf("%d\n", c)
} // 5
// 现在的问题是,使用float会出现精度不够问题,即本来是一个整数,会出现小于这个整数的情况
// 如何处理这类问题?

func consts() {
	// 很奇怪这里使用const常量之后,sqrt()中参数a,b不用强制转换?
	// 这里a,b定义为常量之后类似文本替换,既可以做int也可以做float,所以不用强转
	const a, b = 3, 4
	const filename = "abc.txt"
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func floatPrecision() {
	var a float32 = 135.90
	var c float32 = a * 100
	fmt.Println(c)
}

// 枚举类型直接定义为const即可
func enums() {
	const (
		cpp = iota
		_
		python
		golang
		JavaScript
	)
	// iota 可以作为自增值的一个种子
	//b kb mb gb tb pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, JavaScript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
} // 0 1 2 3
// 1 1024 1048576 1073741824 1099511627776 1125899906842624

func main() {
	fmt.Print("hello world\n")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)
	euler()
	triangle()
	floatPrecision()
	consts()
	enums()
}
