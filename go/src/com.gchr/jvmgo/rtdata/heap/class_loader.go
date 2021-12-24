package heap

import (
	"com.gchr/jvmgo/classpath"
	"com.gchr/jvmgo/classfile"
)

type ClassLoader struct {
	// 通过Classpath来搜索和读取class文件
	cp *classpath.Classpath
	// 记录已经加载的类数据，key是类完全限定名
	classMap map[string]*Class
	// 是否把类加载信息输出控制台
	verboseFlag bool
}

// 创建ClassLoader实例
func newClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	loader := &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
		verboseFlag:verboseFlag,
	}




}

// 加载java.lang.Class类
func (self *ClassLoader) LoadBasicClass() {
	// 加载java.lang.Class类， 会触发java.lang.Object类的加载
	jlClassClass := self.LoadClass("java/lang/Class")
	// 遍历已经加载的每一个类
	for _, class := range self.classMap{
		// 给每个类设置关联对象
		if class.jClass == nil{
			class.jClass = jlClassClass.instanceSlotCount
		}

	}


}

// 把类数据加载到方法区
func (self *ClassLoader) LoadClass(name string) *Class {
	// 判断类是否已经被加载
	if class, ok := self.classMap[name]; ok {
		return class
	}

	var class *Class

	// 若类名的第一个字符是'[', 表示该类是数组类
	if name[0] == '['{
		class = self.loadArrayClass(name)
	} else {
		class = self.loadNonArrayClass(name)
	}
	// 如果已经加载java.lang.class类，则给类关联对象
	if jlClassClass, ok := self.classMap["java/lang/Class"];ok{
		class.jClass = jlClassClass.newObject()
		class.jClass.extra = class
	}
	// 加载类
	return class
}

/*
		数组和普通的类有很大区别
		它的数据并不是来自class文件，而是由Java虚拟机在运行时生成
		暂时不考虑数组类的加载
 */

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	// 找到class文件并把数据读取到内存
	data, entry := self.readClass(name)
}

//读取class
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	// 调用Classpath的ReadClass()方法
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		// 若出错则跑出ClassNotFoundException异常
		panic("java.lang.classNotFoundException: " + name)
	}
	return data, entry
}

// 解析class
func (self *ClassLoader) defineClass(data []byte) *Class {
	// 从byte数组中获取Class结构体
	class := parseClass(data)
	// 设置Class结构它的classloader指针
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

// 把Class文件转换成class结构体
func parseClass(data []byte) *Class {
	// 从byte数组中读取ClassFile
	cf, err := classfile.Parse(data)
	if err != nil {
		panic(err)
	}
	//返回class结构体
	return newClass(cf);
}

// 解析超类符号的引用
func resolveSuperClass(class *Class) {
	// 除了java.lang.Object, 否则要递归调用LoadClass() 方法加载超类
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

//解析接口符号引用
func resolveInterfaces(class *Class) {
	// 获取接口个数
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

// 验证阶段
func verify(class *Class) {
	// todo
}
