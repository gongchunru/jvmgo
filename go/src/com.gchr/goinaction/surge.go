package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {

	ReadConfig("/Users/gongchunru/Documents/newSurge 2.conf")
}

func ReadConfig(path string) error{
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	buf := bufio.NewReader(f)

	//proxy := []byte("[Proxy]")

	for  {
		line, err := buf.ReadString('\n')
		//line, err := buf.Read(proxy)
		if err != nil {
			return err
		}

		fmt.Println(line)
	}
	
	fmt.Println(f)

	return nil

}




