package rtdata

// 帧
type Frame struct {
	// 实现链表数据结构
	lower *Frame
	// 局部变量表指针
	localVars LocalVars
	// 保存操作数栈指针
	operandStack *OperandStack
}

// 实例化帧
/*
	执行方法所需要的局部变量表的大小和操作数栈深度是由编译器预先计算好的
	存储在class文件的method_info结构的Code属性中
 */
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:newLocalVars(maxLocals),
		operandStack:newOperandStack(maxStack),
	}
}
// localVars和operandStackGet方法
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

