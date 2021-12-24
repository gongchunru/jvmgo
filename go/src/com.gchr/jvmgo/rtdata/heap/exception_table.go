package heap

import "com.gchr/jvmgo/classfile"

type ExceptionTable []*ExceptionHandler

type ExceptionHandler struct {
	// tyr{} 语句块的第一条指令
	startPC int

	// try{} 语句块的下一条指令
	endPC     int
	handlerPC int
	//若catchType 为nil 则处理所有异常
	catchType *ClassRef
}

// 创建异常处理表
func newExceptionTable(entries []*classfile.ExceptionTableEntry, cp *classfile.ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPC:   int(entry.StartPc()),
			endPC:     int(entry.EndPc()),
			handlerPC: int(entry.HandlerPc()),
			catchType: getCacheType(uint(entry.CatchType()), cp),
		}
	}
	return table
}

func getCacheType(index uint, cp *ConstantPool) *ClassRef {
	// index为0表示捕获所有异常
	if index == 0 {
		return nil
	}
	return cp.GetConstant(index).(*ClassRef)
}

// 查看异常处理表能否处理异常
func (self ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range self {
		/*
			处理try{}语句块抛出的异常
			try{}语句块包含startPC,但不包含endPC
		 */
		if pc >= handler.startPC && pc < handler.endPC {
			if handler.catchType == nil {
				return handler
			}
			catchClass := handler.catchType.ResolvedClass()
			if catchClass == exClass || catchClass.isSubClassOf(exClass) {
				return handler
			}
		}
	}
	return nil
}
