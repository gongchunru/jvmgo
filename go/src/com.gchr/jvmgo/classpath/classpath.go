package classpath

import (
	"os"
	"path/filepath"
)

//classpath结构体
type Classpath struct {
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}

/*
	Parse函数使用 -Xjre选项解析启动类路径和扩展类路径
	使用-classpath/-cp 选项解析用户路径
 */

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}

	//解析启动类路径和扩展类路径
	cp.parseBootAndExtClasspath(jreOption)
	//解析用户类路径
	cp.parseUserClasspath(jreOption)

	return cp


}
/*
	解析启动类路径和扩张类路径的方法
 */
func (self  *Classpath) parseBootAndExtClasspath (jreOption string)  {
	//获取jre目录
	//jreDir :=
	jreDir := getJreDir(jreOption)

	//加载jre目录下的所有jar包(jreDir/lib/*)
	jreLibPath := filepath.Join(jreDir,"lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	//加载jre目录下所有扩展jar包(jreDir/lib/ext/*)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)

}

/*
解析用户类路径
 */

func (self *Classpath) parseUserClasspath(cpOption string)  {
	//如果用户没有提供-claspath/-cp 选项,则使用当前目录作为用户类路径
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

/*
	获取jre目录
	优先使用用户输入的 -Xjre选项作为jre目录
	如果用户没有输入 -Xjre， 则在当前目录下寻找jre目录
	如果当前目录下没有jre目录， 则尝试使用JAVA_HOME环境变量
 */
func getJreDir(jreOption string) string {
	//判断用户输入的 -Xjre
	if(jreOption != "" && exists(jreOption)){
		return jreOption
	}

	//寻找当前目录下是否存在jre目录
	if exists("./jre") {
		return "./jre"
	}

	//尝试使用JAVA_HOME环境变量
	if jh := os.Getenv("JAVA_HOME"); jh != ""{
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre folder!")
}

//搜索class方法
func (self *Classpath) ReadClass(className string) ([]byte, Entry,error)  {
	//访问ReadClass方法秩序传递类名，不用包含 ".class"后缀
	className = className + ".class"
	//从bootClasspath搜索class文件
	data, entry, err := self.bootClasspath.readClass(className)
	if err == nil {
		return data, entry, err
	}
	
	//从extClasspath 搜索class文件
	data2, entry2, err2 := self.extClasspath.readClass(className)
	if err2 == nil {
		return data2,entry2,err2
	}

	//从userClasspath搜索class文件
	return self.userClasspath.readClass(className)
}

//exists 函数用户判断目录是否存在
func exists(path string) bool {
	_, err := os.Stat(path)
	if err !=nil{
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//返回用户类路径字符串
func (self *Classpath) strings() string  {
	return self.userClasspath.String()
}


