package classfile


type SourceFileAttribute struct {
	cp ConstantPool
	sourceFileIndex uint16
}

// 读取SourceFile_attribute
func (self *SourceFileAttribute) readInfo(reader *ClassReader){
	// attribute_length值为2，所以固定读取2个字节
	self.sourceFileIndex = reader.readUint16()
}

// 通过索引读取常量池中utf8 字面值
func (self *SourceFileAttribute) FileName() string  {
	return self.cp.getUtf8(self.sourceFileIndex)
}