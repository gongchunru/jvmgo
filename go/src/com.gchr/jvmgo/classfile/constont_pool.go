package classfile

import "fmt"

type ConstantPool []ConstantInfo

//读取常量池
func readConstantPool(reader *ClassReader) ConstantPool  {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	for i :=1;  i < cpCount; i++  {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ //占两个位置
		}
	}
	return cp
}

//通过索引读取常量结构体
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo{
	/*
		为什么不是index - 1 ?
		因为常量池的索引时从1 开始的
	 */
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	//panic("Invalid constant pool index!")
	panic(fmt.Errorf("Invalid constant pool index: %v!", index))

}
// 读取ConstantNameAndTypeInfo中name_index和descriptor_index对应的字面量
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	//通过索引读取ConstantNameAndTypeInfo
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name,_type

}
// 从常量池查找类名
func (self ConstantPool) getClassName(index uint16) string {
	//通过索引读取CONSTANT_Class_info
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

// 通过索引获取ConstantUtf8Info结构体的字面值
func (self ConstantPool) getUtf8(index uint16)  string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}