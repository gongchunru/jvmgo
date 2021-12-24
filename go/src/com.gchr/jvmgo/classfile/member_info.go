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
	return self.accessFlags
}
//从常量池中查找字段和方法名
func (self *MemberInfo) Name() string  {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

// 读取code属性
func (self *MemberInfo) CodeAttribute() *CodeAttribute  {
	// 遍历一个method_info中的attribute
	for _,attrInfo := range self.attributes{
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}


// 读取ConstantValue属性
func (self *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute  {
	// 遍历field_info中的attributes
	for _, attrInfo := range self.attributes{
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
return nil
}











