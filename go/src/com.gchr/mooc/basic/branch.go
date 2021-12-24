package main

import (
	"io/ioutil"
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 函数做为一个参数
func apply(op func(int, int) int, a, b int) int {
	// 利用反射获取方法名称
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	// 打印方法名称
	fmt.Printf("Calling function %s with args "+" (%d, %d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
func sum(numbers ...int) int {
	sum := 0
	for i := range numbers{
		sum += numbers[i]
	}
	return sum
}

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(apply(pow, 3, 4))

	// 利用匿名函数

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

}
