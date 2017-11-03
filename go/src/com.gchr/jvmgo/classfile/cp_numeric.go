package classfile

import "math"

/*
	CONSTANT_Integer_info正好可以容纳一个Java的int型常量
	实际上比int更小的boolean、byte、short和char类型的常量也放在
	CONSTANT_Integer_info中
 */

type ConstantIntegerInfo struct {
	val int32
}

//先读取一个uint32数据，然后把它转型成int32类型
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader)  {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

//获取int32的值
func (self *ConstantIntegerInfo) Value() int32  {
	return self.val
}

type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader)  {
	//读取4个字节
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

// 获取float32常量值
func (self *ConstantFloatInfo) Value() float32  {
	return self.val
}

// ConstantLongInfo
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader)  {
	//读取8个字节
	bytes := reader.readUint64()
	//将bytes转换为int64
	self.val = int64(bytes)
}

//获取int64常量值
func (self *ConstantLongInfo) Value() int64  {
	return self.val
}

//ConstantDoubleInfo
type ConstantDoubleInfo struct {
	val float64
}

//读取ConstantDoubleInfo
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader)  {
	//读取8个字节
	bytes := reader.readUint64()
	//将bytes转换为float64
	self.val = math.Float64frombits(bytes)
}

func (self *ConstantDoubleInfo) Value() float64  {
	return self.val
}





