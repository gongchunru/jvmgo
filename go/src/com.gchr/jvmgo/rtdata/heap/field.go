package heap

import "com.gchr/jvmgo/classfile"

// 字段信息结构体
type Field struct {
	// 字段和方法公用结构体
	ClassMember

	// 字段在slot中的位置
	slotId uint
}


// 初始化字段信息
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field{
	fields := make([]*Field, len(cfFields))

	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
	}
	return fields
}





