package classfile

import (
	"fmt"
	"unicode/utf16"
)

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

func (self *ConstantUtf8Info) Str() string  {
	return self.str
}

func decodeMUTF8(bytearr []byte) string {
	utflen := len(bytearr)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count - 1])
			if char2 & 0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c & 0x1F << 6 | char2 & 0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count - 2])
			char3 = uint16(bytearr[count - 1])
			if char2 & 0xC0 != 0x80 || char3 & 0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c & 0x0F << 12 | char2 & 0x3F << 6 | char3 & 0x3F << 0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}