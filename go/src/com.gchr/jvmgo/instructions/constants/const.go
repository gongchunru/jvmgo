package constants

import (
	"com.gchr/jvmgo/instructions/base"
	"com.gchr/jvmgo/rtdata"
)

type ACONST_NULL struct {
	base.NoOperandsInstruction
}

type DCONST_0 struct{ base.NoOperandsInstruction }

// dconst_0 把double类型的0推入操作数栈
func (self *DCONST_0) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushDouble(0.0)
}

// dconst_1 把double类型的1 推入操作数栈
type DCONST_1 struct{ base.NoOperandsInstruction }

func (self *DCONST_1) Execute(frame *rtdata.Frame)  {
	frame.OperandStack().PushDouble(1.0)
}

// 把float类型的0值推入操作数栈
type FCONST_0 struct{ base.NoOperandsInstruction }

func (self *FCONST_0) Execute(frame *rtdata.Frame)  {
	frame.OperandStack().PushFloat(0.0)
}

// 把float类型的1推入操作数栈
type FCONST_1 struct{ base.NoOperandsInstruction }

func (self *FCONST_1) Execute(frame *rtdata.Frame)  {
	frame.OperandStack().PushFloat(1.0)
}


// 把float类型的2推入操作数栈
type FCONST_2 struct{ base.NoOperandsInstruction }

func (self *FCONST_2) Execute(frame *rtdata.Frame)  {
	frame.OperandStack().PushFloat(2.0)
}


//把null引用推入操作数栈顶
func (self *ACONST_NULL) Execute(frame *rtdata.Frame)  {
	frame.OperandStack().PushRef(nil)
}

//把int型-1 推入操作数栈
type ICONST_M1 struct{ base.NoOperandsInstruction }

func (self *ICONST_M1) Execute(frame *rtdata.Frame)  {
	frame.OperandStack().PushInt(-1)
}

// inconst_0 指令把int类型的0推入操作数栈
type ICONST_0 struct{ base.NoOperandsInstruction }
func (self *ICONST_0) Execute(frame *rtdata.Frame)  {
	frame.OperandStack().PushInt(0)
}

// iconst_1 把指令int类型的1推入操作数栈
type ICONST_1 struct{ base.NoOperandsInstruction }

func (self *ICONST_1) Execute(frame *rtdata.Frame)  {
	frame.OperandStack().PushInt(1)
}

// iconst_2 把指令int类型的2推入操作数栈
type ICONST_2 struct{ base.NoOperandsInstruction }

func (self *ICONST_2) Execute(frame *rtdata.Frame)  {
	frame.OperandStack().PushInt(2)
}

// iconst_3 把指令int类型的3推入操作数栈
type ICONST_3 struct{ base.NoOperandsInstruction }

func (self *ICONST_3) Execute(frame *rtdata.Frame)  {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct{ base.NoOperandsInstruction }
func (self *ICONST_4) Execute(frame *rtdata.Frame)  {
	frame.OperandStack().PushInt(4)
}


type ICONST_5 struct{ base.NoOperandsInstruction }
func (self *ICONST_5) Execute(frame *rtdata.Frame)  {
	frame.OperandStack().PushInt(5)
}


type LCONST_0 struct{ base.NoOperandsInstruction }
func (self *LCONST_0) Execute(frame *rtdata.Frame)  {
	frame.OperandStack().PushLong(0)
}

type LCONST_1 struct{ base.NoOperandsInstruction }
func (self *LCONST_1) Execute(frame *rtdata.Frame)  {
	frame.OperandStack().PushLong(1)
}






