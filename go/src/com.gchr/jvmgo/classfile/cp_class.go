package classfile

type ConstantClassInfo struct {
	cp ConstantPool
	nameIndex uint16
}

// 读取CONSTANT_Class_info
func (self *ConstantClassInfo) readInfo(reader *ClassReader)  {
	self.nameIndex = reader.readUint16()
}

// 读取class字面值
func (self *ConstantClassInfo) Name() string  {
	return self.cp.getUtf8(self.nameIndex)
}







