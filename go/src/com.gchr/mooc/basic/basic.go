package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

func euler() {
	c := 3 + 4i
	fmt.Print(cmplx.Abs(c))

	// 欧拉公式  float类型，结果不是0
	fmt.Print(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("\n %.3f\n", cmplx.Exp(1i*math.Pi)+1)

	// 这种有问题的写法
	fmt.Print(cmplx.Pow(math.E, 1i+math.Pi) + 1)

}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Print(c)
}

func consts() {
	const (
		filename = "aba"
		a, b     = 3, 4
	)

	var c int
	c = a * b
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp    = iota
		java
		python
		golang
	)




	fmt.Println(cpp,java,python,golang)

}

func main() {
	euler()
	fmt.Println("")
	triangle()
	fmt.Println("")
	consts()
	fmt.Println("")

	enums()



}
