package heap

/*
	将类名转换成类型描述符
	XXX -> [LXXX;
	int -> [I
	[XXX -> [[XXX
 */

func getArrayClassName(classname string) string {
	return "[" + toDescriptor(classname)
}

/*
	将类名转成类型描述符
	XXX  => LXXX;
	int  => I
	[XXX => [XXX
 */

func toDescriptor(className string) string {
	// 如果已经是数组名，直接返回
	if className[0] == '[' {
		return className
	}

	// 如果是基本类型名，返回对应的类型描述符
	if d, ok := primitiveTypes[className]; ok {
		return d
	}

	// 普通类名转成类型描述符
	return "L" + className + ";"

}

// // 基本类型和对应类型描述符map
var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

/*
	将数组类名转换成类名
	[[XXX -> [XXX
	[LXXX; -> XXX
	[I -> int
 */
func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}

/*
// [XXX  => [XXX
// LXXX; => XXX
// I     => int

 */
func toClassName(descriptor string) string {
	// 若是数组， 直接返回
	if descriptor[0] == '[' {
		return descriptor
	}
	// 如果是引用类型则去掉前缀L
	if descriptor[0] == 'L' {
		return descriptor[1 : len(descriptor)-1]
	}

	// 若是基本类型，通过基本类型描述符获取基本类型
	for className, d := range primitiveTypes {
		if d == descriptor {
			return className
		}
	}

	panic("Invalid descriptor:" + descriptor)
}
