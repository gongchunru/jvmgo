package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"fmt"
	"os/exec"
	"com.gchr/goinaction/ch2/sample/search"
	_ "com.gchr/goinaction/ch2/sample/matchers"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main()  {
	search.Run("president")
	//dir :=getCurrentDirectory()
	//fmt.Println(dir)

	//getCurrentDir()
	//execpath, err := os.Executable()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Println(execpath)

}


func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func getCurrentDir() string  {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Println(err.Error())
	}
	s = strings.Replace(s, "\\", "/", -1)
	s = strings.Replace(s, "\\\\", "/", -1)
	i := strings.LastIndex(s, "/")
	path := string(s[0 : i+1])
	fmt.Println(path)
	return path

}