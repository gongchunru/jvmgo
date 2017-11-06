package rtdata
/*
|-JVM
|--Thread
|---pc
|---Stack
|		Frame
|			LocalVars
|			OperandStack
|

 */

type Thread struct {
	//pc寄存器，存放当前正在执行的Java虚拟机指令的地址
	pc int
	// 虚拟机栈结构体指针
	stack *Stack
}



func NewThread() *Thread{
	return &Thread{
		stack:newStack(1024),
	}
}
func (self *Thread) PC() int  {

}

func (self *Thread) PushFrame(frame *Frame) {
	//self.stack.
}

func (self *Thread) PopFrame() *Frame {
	//return self.stack
}

func (self *Thread) CurrentFrame() *Frame  {
	//return
}
