package base

import "com.gchr/jvmgo/rtdata"

// 定义接口

type Instruction interface {
	//从字节码中取操作数
	FetchOperands(reader *BytecodeReader)
	//执行指令逻辑
	Execute(frame *rtdata.Frame)
}

//表示没有操作数的指令
type NoOperandsInstruction struct {

}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader)  {

}
// 跳转指令
type BranchInstruction struct {
	// 跳转偏移量
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader)  {
	self.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	//表示局部变量表索引
	Index uint
}
// 从字节码中读取一个int8 整数， 转成uint后赋值给Index字段
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader)  {
	self.Index = uint(reader.ReadUint8())
}