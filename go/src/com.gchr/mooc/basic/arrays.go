package main

import "fmt"


/*
go只有值传递，传参就是拷贝数值

[10]int 和[20]int是不同的类型
调用func f(arr [10]int)会拷贝数组
 */

func printArray(arr [5]int)  {
	arr[0] = 100
	for i,v := range arr{
		fmt.Println(i,v)
	}

}


func main() {
	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{1,2,3,4,5}

	var grid [4][5]bool

	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)

	//遍历数组
	for i, v := range arr3{
		fmt.Println(i,v)
	}


	fmt.Println("printArray(arr1)")
	printArray(arr1)

	fmt.Println("printArray(arr3)")
	printArray(arr3)
	//报错
	//printArray(arr2)

	fmt.Println("pring arr1 and arr3")
	fmt.Println(arr1,arr3)
	





}
