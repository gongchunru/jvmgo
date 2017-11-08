package constants

import (
	"com.gchr/jvmgo/instructions/base"
	"com.gchr/jvmgo/rtdata"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame rtdata.Frame)  {
	// 什么也不做
}