package rtdata

// 局部变量表结构体
type LocalVars []Slot

// 创建LocalVars实例
func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot,maxLocals)
	}
	return nil
}
