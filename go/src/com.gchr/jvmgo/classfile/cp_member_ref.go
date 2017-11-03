package classfile

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}


/*
	CONSTANT_Fieldref_info、CONSTANT_Methodref_info和CONSTANT_InterfaceMethodref_info 的结构一样
	所以使用一样的结构体ConstantMemberrefInfo来表示
 */
type ConstantMemberrefInfo struct {
	cp ConstantPool
	classIndex uint16
	nameAndTypeIndex uint16
}

// 读取ConstantMemberrefInfo
func (self *ConstantMemberrefInfo) readInfo (reader *ClassReader)  {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex  = reader.readUint16()
}

// 通过classIndex读取class字面值
func (self *ConstantMemberrefInfo) ClassName()  string {
	return self.cp.getClassName(self.classIndex)
}

// // 通过nameAndTypeIndex读取名称和描述字面值
func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string)  {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}