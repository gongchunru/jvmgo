package constants

import (
	"com.gchr/jvmgo/instructions/base"
	"com.gchr/jvmgo/rtdata"
)

type ACONST_NULL struct {
	base.NoOperandsInstruction
}

type DCONST_0 struct{ base.NoOperandsInstruction }
type DCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_0 struct{ base.NoOperandsInstruction }
type FCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_2 struct{ base.NoOperandsInstruction }
type ICONST_M1 struct{ base.NoOperandsInstruction }
type ICONST_0 struct{ base.NoOperandsInstruction }
type ICONST_1 struct{ base.NoOperandsInstruction }
type ICONST_2 struct{ base.NoOperandsInstruction }
type ICONST_3 struct{ base.NoOperandsInstruction }
type ICONST_4 struct{ base.NoOperandsInstruction }
type ICONST_5 struct{ base.NoOperandsInstruction }
type LCONST_0 struct{ base.NoOperandsInstruction }
type LCONST_1 struct{ base.NoOperandsInstruction }

//把null引用推入操作数栈顶
func (self *ACONST_NULL) Execute(frame rtdata.Frame)  {
	frame.OperandStack().PushRef(nil)
}

// 把double型的0值推入操作数栈
func (self *ICONST_0) Execute(frame rtdata.Frame)  {
	frame.OperandStack().PushDouble(0.0)
}
//把int型-1 推入操作数栈
func (self *ICONST_M1) Execute(frame rtdata.Frame)  {
	frame.OperandStack().PushInt(-1)
}
