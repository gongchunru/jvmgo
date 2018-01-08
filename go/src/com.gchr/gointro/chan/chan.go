package main

import "fmt"

func sum(a []int, c chan interface{})  {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan interface{})

	go sum(a[:len(a)/2],c)

	x, y := <-c, <-c

	y1,ok2 := y.(int)
	if x1, ok1 := x.(int); ok1 && ok2{
		fmt.Println(x,y,x1 + y1)
	}

}
