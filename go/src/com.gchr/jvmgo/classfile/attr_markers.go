package classfile


// 统一的结构体
type MarkerAttribute struct{}


type DeprecatedAttribute struct {
	MarkerAttribute
}



type SyntheticAttribute struct {
	MarkerAttribute
}

// attribute_length 值为0， 所以方法中什么都不做
func (self *MarkerAttribute) readInfo(reader *ClassReader)  {

}


