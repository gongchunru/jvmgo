package constants

import "com.gchr/jvmgo/instructions/base"

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



