package classfile

type AttributeInfo interface {
	// 读取属性信息由具体的属性结构实现
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	// 读取2个字节的属性长度
	attributesCount := reader.readUint16()
	//读取所需属性
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes  {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	//读取2个字节的attribute_name_index
	attrNameIndex := reader.readUint16()
	//通过attrNameIndex从常量池中获取utf8字面值
	attrName := cp.getUtf8(attrNameIndex)
	//读取固定4个字节的attribute_length
	attrLen := reader.readUint32()

	//读取属性
	attrInfo := newAttributeInfo(attrName, attrLen, cp)

	attrInfo.readInfo(reader)
	return attrInfo
}

//通过属性名返回属性实现
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		// 读取未实现的属性
		return &UnparsedAttribute{attrName, attrLen, nil}
	}

}
