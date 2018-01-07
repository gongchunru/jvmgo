package main

import (
	"net/http"
	"fmt"
	"time"
)

type Hello struct {}

func (h Hello) ServerHttp( w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello!")
}


func main() {

	//http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
	//	fmt.Fprint(w,"hello world")
	//
	//})
	//err := http.ListenAndServe("localhost:4000",nil)
	//if err != nil{
	//	log.Fatal(err)
	//}

	go say("hello")
	say("world")

}

func say(s string)  {
	for i:=0 ;i < 5 ; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d : %s \n",i,s)
	}
}



