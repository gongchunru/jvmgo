package heap

// Object 结构体

type Object struct {
	// 对象的Class指针
	//class *Class

	// 实例变量，可以容纳任何类型的值
	data interface{}
	extra interface{}

}
