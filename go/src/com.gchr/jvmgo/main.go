package main


import (
	"com.gchr/jvmgo/cmd"
	 "fmt"

	"com.gchr/jvmgo/classpath"
	"strings"
	"com.gchr/jvmgo/classfile"
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

	// 将.替换成/ （java.lang.String -> java/lang/String）
	className := strings.Replace(cmd.Class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(cmd.Class)
	printClassInfo(cf)
}

// 读取并解析class文件
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

//把class文件的一些重要信息打印出来
func printClassInfo(cf *classfile.ClassFile)  {
	fmt.Printf("Version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count :%v\n",len(cf.ConstantPool()))
	fmt.Printf("access flags : 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf(" %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf(" %s\n", m.Name())
	}
}
