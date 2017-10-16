package classpath

import (
	"path/filepath"
	"io/ioutil"
)

//表示目录形式的类路径

type DirEntry struct {
	//用于存放绝对路径
	absDir string
}

/*
	返回执行DirEntry对象的指针
	GO语言没有专门的构造函数，此函数就当做DirEntry的构造函数
 */
func NewDirEntry(path string) *DirEntry{
	/*
		GO使用error值来表示错误的状态
		Go使用panic和recover来处理所谓的运行时异常
	 */

	absDir, err := filepath.Abs(path) //把参数转为绝对路径
	if err != nil {
		panic(err) //如果出错，则终止执行
	}
	return &DirEntry{absDir}
}

/*
	Go没有类，可以使用方法接受者的方式在结构体类型上定义方法
	指向DirEntry对象的指针self为方法接受者
	该方法用来读取class文件文件
 */
func (self *DirEntry) readClass(className string)([]byte, Entry, error)  {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

//实现String方法
func (self *DirEntry) String() string {
	return self.absDir
}


