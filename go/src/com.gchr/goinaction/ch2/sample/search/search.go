package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)



func Run(searchTerm string)  {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	for _, feed := range feeds{
		matcher, exists := matchers[feed.Type]

		if !exists{
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个goroutine 来监控是否所有工作都做完
	go func(){
		// 等候所有的任务做完
		waitGroup.Wait()

		// 可以关闭通道的方式，通知display函数
		// 可以退出程序了
		close(results)
	}()

	// 启动函数，显示返回结果
	// 并且在最后一个结果显示完后返回
	Display(results)
}

// Register调用时，会注册一个匹配器，提供给后面的程序使用
// 将一个 Matcher 值加入到保存注册匹配器的映射中。所有这种注 册都应该在 main 函数被调用前完成
func Register(feedType string, matcher Matcher)  {
	if _, ok := matchers[feedType]; ok{
		log.Fatalln(feedType,"Matcher already registered")
	}

	log.Println("Register", feedType,"matcher")
	matchers[feedType] = matcher

}





