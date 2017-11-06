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
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:newLocalVars(maxLocals),
		//operandStack:
	}
}