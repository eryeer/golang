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
	fmt.Printf("%d %s\n", a, s)
}

func variableDeduction() {
	var a, b, c, d = 3, 5, true, "aaa"
	fmt.Println(a, b, c, d)
}
func variableShort() {
	a, b, c, d := 3, 5, true, "aaa"
	fmt.Println(a, b, c, d)
}

func main() {
	fmt.Println("hello world")
	//variableInit()
	variableDeduction()
	fmt.Println(aaa,bbb,ccc)
}
