package heap

import "com.gchr/jvmgo/classfile"

type Method struct {
	//字段和方法公用结构体
	ClassMember

	//操作数栈最大深度
	maxStack uint

	// 局部变量表大小
	maxLocals uint
	// 字节码
	code []byte
	// 方法的参数所占的Slot数量
	argSlotCount uint
	// 异常处理表
	exceptionTable ExceptionTable
	//行号表信息
	lineNumberTable *classfile.LineNumberTableAttribute
}

//初始化方法信息
func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method  {
	methods := make([]*Method,len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class,cfMethod)
	}
	return methods
}

func newMethod(class *Class,cfMethod *classfile.MemberInfo) *Method  {
	method := &Method{}
	method.class = class
	//method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	return method
}

// 从方法中读取maxStack,maxLocals和字节码
func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo)  {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil{
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
		self.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(), self.class.constantPool)
		self.lineNumberTable = codeAttr.LineNumberTableAttribute()
	}
}

// 查找异常处理表
func (self *Method) FindExceptionHandler(exClass *Class, pc int) int  {
	handler := self.exceptionTable
}









