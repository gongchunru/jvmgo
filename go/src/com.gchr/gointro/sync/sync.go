package main

import (
	"sync"
	"runtime"
	"fmt"
	"sync/atomic"
)

var (
	count int32
	wg sync.WaitGroup
)

func incCount()  {
	defer wg.Done()
	for i:=0; i < 2; i++{
		//value := count
		value := atomic.LoadInt32(&count)
		runtime.Gosched()
		value ++
		//count = value
		atomic.StoreInt32(&count,value)
	}

}

func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}
