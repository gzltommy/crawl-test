package main

import (
	"github.com/gzltommy/crawl-test/04.MulTask/engine"
	"github.com/gzltommy/crawl-test/04.MulTask/parse"
	"github.com/gzltommy/crawl-test/04.MulTask/scheduler"
	"github.com/gzltommy/crawl-test/04.MulTask/types"
)

func main() {
	// 最终：从最开始爬，并嵌套爬
	e := engine.ConcurrentEngine{
		WorkCount: 100,
		Scheduler: &scheduler.SimpleScheduler{},
	}

	e.Run(types.Request{
		Url:      "https://book.douban.com/",
		ParseFun: parse.Tag,
	})
}
