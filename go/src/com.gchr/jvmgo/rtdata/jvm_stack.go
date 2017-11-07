package rtdata

// 虚拟机栈结构体
type Stack struct {
	// 栈的容量 （最多可以容纳多少帧）
	maxSize uint
	// 当前栈大小
	size uint
	// 栈顶指针
	_top  *Frame
}

func newStack(maxSize uint) *Stack  {
	return &Stack{
		maxSize:maxSize,
	}
}
// 入栈
func (self *Stack) push(frame *Frame)  {
	// 超过栈最大深度报错
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}

	// 将栈顶指针下移 todo why?
	if self._top != nil {
		frame.lower = self._top
	}

	self._top = frame
	self.size++
}

// 出栈 todo thinking why
func (self *Stack) pop() *Frame  {
	if self._top == nil{
		panic("jvm stack is empty!")
	}
	top := self._top
	// 弹出栈顶
	self._top = top.lower
	// 弹出指针的lower指向不再有意义
	top.lower = nil
	self.size--
	return top
}

// 只返回栈顶指针，不弹出
func (self *Stack) top() *Frame  {
	if self._top == nil{
		panic("jvm stack is empty@")
	}
	return self._top
}


