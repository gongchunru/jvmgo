package rtdata

import "com.gchr/jvmgo/rtdata/heap"

// 帧
type Frame struct {
	// 实现链表数据结构
	lower *Frame
	// 局部变量表指针
	localVars LocalVars
	// 保存操作数栈指针
	operandStack *OperandStack
	//线程指针
	thread *Thread

	// 下个PC寄存器地址（为了实现跳转）
	nextPC int

	//方法区的方法信息指针
	method *heap.Method
}

// NewFrame 实例化帧
/*
	执行方法所需要的局部变量表的大小和操作数栈深度是由编译器预先计算好的
	存储在class文件的method_info结构的Code属性中
*/
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// LocalVars localVars和operandStackGet方法
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

// RevertNextPC 重置nextPC
func (self *Frame) RevertNextPC() {
	// Thread的pc寄存器字段始终指向当前指令地址
	self.nextPC = self.thread.pc
}
