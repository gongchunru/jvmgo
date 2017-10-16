package cmd

import (
	"fmt"
	"flag"
	"os"
)

/*
定义结构体，用于存储命令参数
java [-options] class [args...]
 */
type Cmd struct {
	HelpFlag bool
	VersionFlag bool
	CpOption string
	XjreOption string
	Class  string
	Args []string
}

/*
ParseCmd()方法返回值为*Cmd, 是指向Cmd的值的指针
Go语言中，常量、函数的首字母大写表示对外公开的相当于Java的public，小写的表示私有的相当于Java的private
 */
func ParseCmd() *Cmd  {
	//声明cmd为指向空的Cmd对象的指针
	cmd := &Cmd{}

	/*
	Usage是一个函数，默认输出所有定义了的命令行参数和帮助信息

	 */

	 flag.Usage = PrintUsage
	 //flag.XxxVar(),将flag绑定到第一个参数指定的指针上，并设置默认值
	 flag.BoolVar(&cmd.HelpFlag,"help",false,"print help message")
	 flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
	 flag.BoolVar(&cmd.VersionFlag, "version", false,"print version and exit")
	 flag.StringVar(&cmd.CpOption, "classpath","","classpath")
	 flag.StringVar(&cmd.CpOption,"cp","","classpath")
	 flag.StringVar(&cmd.XjreOption,"Xjre","","path to jre")
	 //在所有的flag定义完成之后，可以通过调用flag.Parse()进行解析
	 flag.Parse()

	 args := flag.Args()
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}

	return cmd
}

func PrintUsage()  {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])

}

