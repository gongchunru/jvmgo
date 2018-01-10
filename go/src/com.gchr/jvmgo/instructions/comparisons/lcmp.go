package comparisons

import (
	"com.gchr/jvmgo/instructions/base"
	"com.gchr/jvmgo/rtdata"
)

/*
用于比较lang变量，
 */
// Long 比较指令结构体
type LCMP struct {
	base.NoOperandsInstruction
}

func (self *LCMP) Execute(frame *rtdata.Frame)  {
	// 弹出栈顶long型变量，作为比较符后面的操作数
	stack := frame.OperandStack()

	//弹出栈顶long型变量，作为比较符前面的操作数
	v2 := stack.PopLong()

	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
