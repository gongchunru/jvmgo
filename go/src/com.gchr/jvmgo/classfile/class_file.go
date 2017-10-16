package classfile

type ClassFile struct {
	magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields []*MemberInfo
	methods []*MemberInfo
	attributes []AttributeInfo
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader){
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		/*
			Java虚拟机的实现是抛出java.lang.ClassFormatError
			目前由于才开始写虚拟机，还没法抛出异常，暂先调用panic
		 */
		panic("java.lang.ClassFormatError:magic!")
	}
}
