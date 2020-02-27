package main

import "fmt"

func variableZeroValue() {
	var a int
	var b int
	fmt.Printf("%d,%q\n", a, b)
}

// hello world
// 0,'\x00' 使用%q打印出空串
func main() {
	fmt.Print("hello world\n")

	variableZeroValue()
}
