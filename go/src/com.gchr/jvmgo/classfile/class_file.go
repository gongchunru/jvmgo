package classfile

import (
	"fmt"
	"math/cmplx"
)

/*
	该文件反应的是Java虚拟机规范定义的class文件格式。
	go语言中所有首字母大写的类型、结构体、字段、变量、函数、方法等都是公开的
	小写的只能本包使用
 */
type ClassFile struct {
	magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields []*MemberInfo
	methods []*MemberInfo
	attributes []AttributeInfo
}

// 读取魔数
func (self *ClassFile) readAndCheckMagic(reader *ClassReader){
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		/*
			Java虚拟机的实现是抛出java.lang.ClassFormatError
			目前由于才开始写虚拟机，还没法抛出异常，暂先调用panic
		 */
		panic("java.lang.ClassFormatError:magic!")
	}
}
//读取版本号
func (self *ClassFile) readAndCheckVersion(reader *ClassReader)  {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
		//主版本号是45直接返回
	case 45:
		//主版本号为46-52时，次版本号必须为0
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0{
			return
		}
		return
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// 读取类访问标志
func (self *ClassFile) readAccessFlags(reader *ClassReader)  {
	self.accessFlags = reader.readUint16()
}

// 把[]byte解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile, err error)  {
	defer func() {
		if r := recover(); r != nil{
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr) //read()方法
	return
}

func (self *ClassFile) read(reader *ClassReader){
	self.readAndCheckMagic(reader)

	//self.reade

	self.constantPool = readConstantPool(reader)

	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)

	self.methods = readMembers(reader, self.constantPool)
	//self.attributes = readAttributes()
}
//把结构体的字段暴露 给其他包使用 类似Java中的Get方法
func (self *ClassFile) MinorVersion() uint16  {
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}



//从常量池查找类名
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName()  string {
	if self.superClass > 0{
		return self.constantPool.getClassName(self.superClass)
	}
	return "" //只有 java.lang.Object没有超类
}
// 从常量池中查找接口名
func (self *ClassFile) InterfaceNames() []string  {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces{
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

