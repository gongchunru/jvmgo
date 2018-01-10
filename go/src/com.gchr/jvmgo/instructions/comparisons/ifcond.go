package comparisons

import (
	"com.gchr/jvmgo/instructions/base"
	"com.gchr/jvmgo/rtdata"
)

/*
if<cond>指令
if<cond>指令把操作数栈顶的int变量弹出，然后跟0进行比较，满足条件则跳转。
假设从栈顶弹出的变量时x，则指令执行跳转操作的条件如下：

ifeq: x == 0
ifne: x != 0
iflt: x < 0
ifle: x <= 0
ifgt: x > 0
ifge: x >= 0

 */

 //ifeq
type IFEQ struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *rtdata.Frame)  {
	//弹出栈顶int变量
	val := frame.OperandStack().PopInt()
	if val == 0{
		//满足条件跳转
		base.Branch(frame,self.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (self *IFNE) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (self *IFLT) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (self *IFLE) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (self *IFGT) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (self *IFGE) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}

