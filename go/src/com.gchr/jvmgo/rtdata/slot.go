package rtdata

import "com.gchr/jvmgo/rtdata/heap"

// 用它来实现局部变量表

type Slot struct{
	// 整数
	num int32
	// 引用
	ref *heap.Object
}
