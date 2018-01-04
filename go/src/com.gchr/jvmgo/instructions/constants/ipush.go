package constants

import (
	"com.gchr/jvmgo/instructions/base"
	"com.gchr/jvmgo/rtdata"
)

// bipush 指令从操作数中获取一个byte
// push byte
type BIPUSH struct {
	val int8
}

// push short
type SIPUSH struct {
	val int16
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader)  {
	self.val = reader.ReadInt8()
}

func (self *SIPUSH) Execute(frame *rtdata.Frame)  {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}



