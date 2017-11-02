package classfile

type ConstantUtf8Info struct {
	str string
}
// 先读出 []byte  调用 decodeMUTF8() 解码成Go 字符串
func (self *ConstantUtf8Info) readInfo(reader *ClassReader)  {
	//读取length(2个字节)，并转换成uint32
	length := uint32(reader.readUint16())
	//指定指定长度的字节
	bytes := reader.readBytes(length)
	//将字节转换成MUFT-8
	self.str = decodeMUTF8(bytes)
}

func (self *ConstantUtf8Info) str() string  {
	return self.str()
}

// 改方法是作者参照java.io.DataInputStream.readUTF(DataInpu)写的
func decodeMUTF8(bytes []byte) string  {
	return string(bytes)
}