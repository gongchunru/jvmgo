package rtdata

import (
	"math"
	"com.gchr/jvmgo/rtdata/heap"
)

type OperandStack struct {
	// 记录栈顶位置
	size uint
	// 操作数栈的大小时编译器已经确定的，所以可以用[]Slot实现
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0{
		return &OperandStack{
			slots:make([]Slot,maxStack),
		}
	}
	return nil
}
func (self *OperandStack) PushInt(val int32)  {
	self.slots[self.size].num = val
	self.size++
}

func (self *OperandStack) PopInt() int32  {
	self.size--
	return self.slots[self.size].num
}

func (self *OperandStack) PushFloat(val float32)  {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}

func (self *OperandStack) PopFloat() float32  {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}
// 把long型推入栈顶时，需要拆成两个int变量，弹出时，先弹出两个int变量，然后组装成一个long型
// todo
func (self *OperandStack) PushLong(val int64)  {
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val >> 32)
	self.size += 2
}

func (self *OperandStack) PopLong() int64  {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)
	return int64(high<<32) | int64(low)
}

// double先转long型再按long变量处理
func (self *OperandStack) PushDouble(val float64)  {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

func (self *OperandStack) PopDouble() float64  {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

// 应用类型
func (self *OperandStack) PushRef(ref *heap.Object)  {
	self.slots[self.size].ref = ref
	self.size++
}
// 把Slot结构体的ref字段设置成nil， 为了帮助go的垃圾收集器回收Object结构体实例
func (self *OperandStack) PopRef() *heap.Object {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil
	return ref
}

//推入 ，弹出一个slot
func (self *OperandStack) PushSlot(slot Slot)  {
	self.slots[self.size] = slot;
	self.size++
}
func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}


