package search

import (
	"os"
	"encoding/json"
)

const dataFile = "data.json"

// 引号内的部分叫标记(tag)，
type Feed struct {
	Name string `json:site`
	URI  string `json:link`
	Type string `json:type`
}

/*
读数据文件，返回值是两个参数：
1. 切片
2. error 类型
 */
func RetrieveFeeds() ([]*Feed, error) {
	//打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// 当函数返回时 关闭文件
	defer file.Close()

	// 将文件解码到一个切片里
	// 这个切片的每一项指向一个Feed类型的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// 这个函数不需要检查错误，调用者会做这件事
	return feeds, err
}
