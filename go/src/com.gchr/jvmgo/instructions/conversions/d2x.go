package conversions

import (
	"com.gchr/jvmgo/instructions/base"
	"com.gchr/jvmgo/rtdata"
)

// 将double强转float结构体
type D2F struct {
	base.NoOperandsInstruction
}

func (self *D2F) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()

	//弹出栈顶double型变量
	d := stack.PopDouble()
	// 强转成float
	f := float32(d)
	// 入栈
	stack.PushFloat(f)
}


// 将double强转int
type D2I struct {
	base.NoOperandsInstruction
}

func (self *D2I) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}