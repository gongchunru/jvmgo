package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo  {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members  {
		members[i] = readMember(reader, cp)
	}

	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo  {

	return &MemberInfo{
		cp: cp,
		accessFlags:reader.readUint16(),
		nameIndex:reader.readUint16(),
		descriptorIndex:reader.readUint16(),
		attributes:readAttributes(reader, cp), // 3.4
	}

}

func (self *MemberInfo) AccessFlags() uint16{

}
//从常量池中查找字段和方法名
func (self *MemberInfo) Name() string  {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}