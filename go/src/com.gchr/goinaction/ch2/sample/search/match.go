package search

import (
	"log"
	"fmt"
)

type Result struct {
	Field string
	Content string
}


type Matcher interface {
	Search(feed *Feed,searchTerm string) ([]*Result,error)
}

// Match函数，为每个数据源单独启动goroutine来执行这个函数
// 并发地执行搜索
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result)  {

	// 对特定的匹配器执行搜索
	searchResults, err := matcher.Search(feed,searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 将结果写入通道
	for _, result := range searchResults{
		results <- result
	}
}

// 从每个单独的goroutine接收到结果后 在终端输出
func Display(results chan *Result)  {
	// 通道会一直被阻塞，直到有结果写入，一旦通道被关闭，for循环就会终止
	for result := range results{
		fmt.Println(" %s:\n%s\n",result.Field,result.Content)
	}
}

