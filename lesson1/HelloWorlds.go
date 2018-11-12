package main

import "fmt"

var (
	aaa = 1
	bbb = true
	ccc = "hello"
)

func variableInit() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s) //%q打印引号
}

func variableTypeDeduction() {
	var a, b, c, d = 3, 5, true, "aaa"
	fmt.Println(a, b, c, d)
}
func variableShort() {
	a, b, c, d := 3, 5, true, "aaa" //使用冒号等于赋值变量
	fmt.Println(a, b, c, d)
}

func main() {
	fmt.Println("hello world")
	variableInit()
	variableZeroValue()
	variableTypeDeduction()
	variableShort()
	fmt.Println(aaa, bbb, ccc)
}
