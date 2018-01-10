package comparisons

import (
	"com.gchr/jvmgo/instructions/base"
	"com.gchr/jvmgo/rtdata"
)

/*
if_acmpeq和if_acmpne指令把栈顶的两个引用弹出，根据引用是否相同进行跳转
*/

type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ICMPNE) Execute(frame *rtdata.Frame)  {
	if _acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

func _acmp(frame *rtdata.Frame)  bool{
	stack := frame.OperandStack()
	//弹出栈顶引用
	ref2 := stack.PopRef()
	//弹出栈顶引用
	ref1 := stack.PopRef()
	//判断是否相等
	return ref1 == ref2
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (self *IF_ICMPNE) Execute(frame *rtdata.Frame)  {
	if !_acmp(frame){
		base.Branch(frame,self.Offset)
	}
}


