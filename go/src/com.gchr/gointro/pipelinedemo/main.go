package main

import (
	"com.gchr/gointro/pipeline"
	"fmt"
	"os"
)

func main() {
	const filename  = "small.in"
	const n  = 50


	file , err :=os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n)
	pipeline.WriterSink(file,p)

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p =  pipeline.ReaderSource(file,-1)
	for v := range p {
		fmt.Println(v)
	}


}


func mergetDemo() {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(
			pipeline.ArraySource(12, 2, 8, 5, 9)))
	//第一种写法
	//for  {
	//	if num, ok := <-p ; ok{
	//		fmt.Println(num)
	//	}else {
	//		break
	//	}
	//}
	// 简便写法：
	for v := range p {
		fmt.Println(v)
	}
}
