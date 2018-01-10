package comparisons

import (
	"com.gchr/jvmgo/instructions/base"
	"com.gchr/jvmgo/rtdata"
)

/*
fcmpg
fcmpl 
用于比较float变量。由于浮点数计算可能产生NaN值
由于浮点数计算有可能产生NaN值，所以比较两个浮点数时除了大于、等于、小于之外，
还有第4种结果：无法比较。
fcmpg和fcmpl指令的区别就在于对第4种结果的定义

 */

type FCMPG struct {
	base.NoOperandsInstruction
}

func (self *FCMPG) Execute(frame *rtdata.Frame)  {
	_fcmp(frame,true)
}


type FCMPL struct {
	base.NoOperandsInstruction
}

func (self *FCMPL) Execute(frame *rtdata.Frame)  {
	_fcmp(frame,false)
}
func _fcmp(frame *rtdata.Frame, gFlag bool)  {
	stack := frame.OperandStack()
	//弹出栈顶float型变量，作为比较后面的操作数
	v2 := stack.PopFloat()
	//弹出栈顶long型变量，作为比较前面的操作数
	v1 := stack.PopFloat()

	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if gFlag {
		// fcmpg指令对于第四种结果入栈1
	} else {
		// fcmpl 指令对于第四种结果入栈0
		stack.PushInt(-1)
	}
}
