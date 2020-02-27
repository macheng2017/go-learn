package main

import "fmt"

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
	a, b, c, d := 3, 4, true, "def"
	// 第一次定义变量可以使用 :=
	b = 5
	fmt.Println(a, b, c, d)
}

func main() {
	fmt.Print("hello world\n")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
}
