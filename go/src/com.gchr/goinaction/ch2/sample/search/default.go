package search

// 空结构体
// 创建实例时候不会分配内存
//
type defaultMatcher struct {
}

func init()  {
	var matcher defaultMatcher
	Register("default",matcher)
}


// 这是一个方法
// 我们可以使用 defaultMatcher 类型的值或者指向这个类型值的指针来调用 Search 方 法
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error)  {
	return nil,nil
}
