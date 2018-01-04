package stack

import (
	"com.gchr/jvmgo/instructions/base"
	"com.gchr/jvmgo/rtdata"
)

type POP struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()

	//弹出栈顶变量
	stack.PopSlot()
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP2) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()

	//弹出栈顶变量
	stack.PopSlot()
	stack.PopSlot()
}
