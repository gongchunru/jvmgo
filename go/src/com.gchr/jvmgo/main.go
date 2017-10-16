package main


import (
	"com.gchr/jvmgo/cmd"
	 "fmt"

	"com.gchr/jvmgo/classpath"
	"strings"
)

func main()  {
	command := cmd.ParseCmd()
	if command.VersionFlag {
		fmt.Println("version 0.0.1")
	}else if command.HelpFlag || command.Class == "" {
		cmd.PrintUsage()
	}else {
		startJVM(command)
	}

}

//模拟启动JVM
func startJVM(cmd *cmd.Cmd)  {
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	fmt.Println("classpath:%v class:%v args:%v\n", cp, cmd.Class, cmd.Args)

	// 将.替换成/ （java.lang.String -> java/lang/String）
	className := strings.Replace(cmd.Class, ".", "/", -1)
	classdata, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Println("Could not find or load main class %s\n", cmd.Class)
		return
	}

	fmt.Println("class data:%v\n",classdata)


	//fmt.Printf("classpath:%s class:%s args:%s",cmd.CpOption,cmd.Class, cmd.Args)
}