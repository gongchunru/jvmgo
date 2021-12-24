package heap

import "com.gchr/jvmgo/classfile"

// 符号引用结构体
type ClassRef struct {
	SymRef
}


// 初始化符号引用结构体
func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef  {
	ref := &classInfo

}



